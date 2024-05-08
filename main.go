package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	controller2 "techiebutler/controller"
	"techiebutler/models"
	"techiebutler/repo"
	router2 "techiebutler/router"
)

func main() {
	memoryStore := repo.NewMemoryStore()
	InitData(memoryStore)
	controller := controller2.NewController(memoryStore)
	router := router2.NewRouter(controller)
	router.InitRoutes()
	err := http.ListenAndServe(":8000", router.Engine)
	if err != nil {
		log.Fatal(err)
	}
}

func InitData(repo repo.IRepo) {
	ctx := context.Background()
	empData := []*models.Employee{
		{
			Name:     "Pavan",
			Position: "Developer",
			Salary:   10000,
		},
		{
			Name:     "Valen",
			Position: "Developer 2",
			Salary:   20000,
		},
		{
			Name:     "Lester",
			Position: "Developer",
			Salary:   10000,
		},
	}
	for _, employee := range empData {
		go repo.CreateEmployee(ctx, employee)
	}

	employee := &models.Employee{
		Name:     "Gauresh",
		Position: "Developer",
		Salary:   9025.32,
	}
	id1, _ := repo.CreateEmployee(ctx, employee)
	employee.Position = "Developer 2"
	employee.Salary = 19000
	repo.UpdateEmployee(ctx, id1, employee)

	employee, _ = repo.GetEmployeeByID(ctx, id1)
	fmt.Println(employee)

	employee = &models.Employee{
		Name:     "Hemraj",
		Position: "Developer 3",
		Salary:   30025.32,
	}
	id2, _ := repo.CreateEmployee(ctx, employee)
	repo.DeleteEmployee(ctx, id2)

}
