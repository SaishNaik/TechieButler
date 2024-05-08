package controller

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"techiebutler/models"
	"techiebutler/repo"
)

type Controller struct {
	repo repo.IRepo
}

func NewController(iRepo repo.IRepo) *Controller {
	return &Controller{repo: iRepo}
}

func (c *Controller) GetEmployeeRecords(ctx *gin.Context) {
	// get request object
	req := ctx.Request
	reqCtx := req.Context()
	rw := ctx.Writer
	page, err := strconv.Atoi(req.URL.Query().Get("page"))
	if err != nil || page < 1 {
		err = errors.New("page parameter malformed")
		WriteResponse(rw, nil, err)
		return
	}
	count, err := strconv.Atoi(req.URL.Query().Get("count"))
	if err != nil || count < 1 {
		err = errors.New("count parameter malformed")
		WriteResponse(rw, nil, err)
		return
	}

	ch1 := make(chan *models.GetEmployeeRecordsDTO)
	ch2 := make(chan *models.GetEmployeeCountDTO)
	go func() {
		records, err := c.repo.GetEmployees(reqCtx, count, page)
		ch1 <- &models.GetEmployeeRecordsDTO{
			Records: records,
			Error:   err,
		}
	}()
	go func() {
		count, err := c.repo.GetTotalEmployeeCount(reqCtx)
		ch2 <- &models.GetEmployeeCountDTO{
			Count: count,
			Error: err,
		}
	}()

	result1 := <-ch1
	if result1.Error != nil {
		WriteResponse(rw, nil, err)
		return
	}
	records := result1.Records

	result2 := <-ch2
	if result2.Error != nil {
		WriteResponse(rw, nil, err)
		return
	}
	total := result2.Count

	result := &models.GetEmployeesResult{
		Employees: records,
		Page:      page,
		Count:     count,
		Total:     total,
	}
	WriteResponse(rw, result, nil)
}

func WriteResponse(writer http.ResponseWriter, data interface{}, err error) {
	response := models.Response{}
	if err != nil {
		response.Success = false
		response.Result = nil
		response.Failure = err.Error()
		writer.WriteHeader(500)
	} else {
		writer.WriteHeader(200)
		response.Success = true
		response.Result = data
	}
	writer.Header().Set("Content-Type", "application/json")
	jsonResult, _ := json.Marshal(response)
	_, _ = writer.Write(jsonResult)
}
