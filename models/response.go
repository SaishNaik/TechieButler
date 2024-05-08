package models

type GetEmployeesResult struct {
	Employees []*Employee `json:"employees"`
	Page      int         `json:"page"`
	Count     int         `json:"count"`
	Total     int         `json:"total"`
}

type Response struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result,omitempty"`
	Failure string      `json:"failure"`
}
