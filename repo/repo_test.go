package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/assert/v2"
	"techiebutler/models"
	"testing"
)

func TestRepo(t *testing.T) {
	ctx := context.Background()

	t.Run("Test Create Employee", func(t *testing.T) {
		t.Run("Create Subscription should be create a new record with specified values", func(t *testing.T) {
			ms := NewMemoryStore()
			empData := &models.Employee{
				Name:     "Pavan",
				Position: "Developer",
				Salary:   11000,
			}
			id, err := ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, nil)
			empData.ID = id
			empRecord, err := ms.GetEmployeeByID(ctx, id)
			assert.Equal(t, err, nil)
			assert.Equal(t, empRecord, empData)
		})

		t.Run("Create Subscription should be throw a error if salary,position and name are empty or incorrect", func(t *testing.T) {
			ms := NewMemoryStore()
			empData := &models.Employee{
				Name:     "Pavan",
				Position: "Developer",
				Salary:   11000,
			}
			empData.Salary = 0.0
			id, err := ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, errors.New("name,salary or position is empty or incorrect"))
			assert.Equal(t, id, 0)
			empData.Position = ""
			id, err = ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, errors.New("name,salary or position is empty or incorrect"))
			empData.Name = ""
			id, err = ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, errors.New("name,salary or position is empty or incorrect"))
		})
	})

	t.Run("Test Get Employee By ID", func(t *testing.T) {
		t.Run("Should return the employee record successfully when the id is present", func(t *testing.T) {
			ms := NewMemoryStore()
			empData := &models.Employee{
				Name:     "Pavan",
				Position: "Developer",
				Salary:   11000,
			}
			id, err := ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, nil)
			empData.ID = id
			empRecord, err := ms.GetEmployeeByID(ctx, id)
			assert.Equal(t, err, nil)
			assert.Equal(t, empRecord, empData)
		})
		t.Run("Should return error when the id is not present", func(t *testing.T) {
			ms := NewMemoryStore()
			id := 111111
			empRecord, err := ms.GetEmployeeByID(ctx, id)
			assert.Equal(t, err, errors.New(fmt.Sprintf("GetEmployeeByID: No employee found with id %d", id)))
			assert.Equal(t, empRecord, nil)
		})
	})

	t.Run("Test Update Employee By ID", func(t *testing.T) {
		t.Run("Should update the employee record successfully when the id is present", func(t *testing.T) {
			ms := NewMemoryStore()
			empData := &models.Employee{
				Name:     "Pavan",
				Position: "Developer",
				Salary:   11000,
			}
			id, err := ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, nil)
			empData.ID = id
			empData.Name = "PavanKumar"
			empData.Salary = 20000
			empData.Position = "Developer 2"
			err = ms.UpdateEmployee(ctx, id, empData)
			assert.Equal(t, err, nil)
			empRecord, err := ms.GetEmployeeByID(ctx, id)
			assert.Equal(t, err, nil)
			assert.Equal(t, empRecord, empData)
		})

		t.Run("Should not update the respective field of employee record when salary, position or name is empty or incorrect", func(t *testing.T) {
			ms := NewMemoryStore()
			empData := &models.Employee{
				Name:     "Pavan",
				Position: "Developer",
				Salary:   11000,
			}
			id, err := ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, nil)

			oldName := empData.Name

			updatedEmpData := new(models.Employee)
			updatedEmpData.Name = ""
			updatedEmpData.Salary = 20000
			updatedEmpData.Position = "Developer 2"
			oldPosition := updatedEmpData.Position
			err = ms.UpdateEmployee(ctx, id, updatedEmpData)
			assert.Equal(t, err, nil)
			empRecord, err := ms.GetEmployeeByID(ctx, id)
			assert.Equal(t, err, nil)
			assert.Equal(t, empRecord.Name, oldName)

			updatedEmpData.Position = " "
			err = ms.UpdateEmployee(ctx, id, updatedEmpData)
			assert.Equal(t, err, nil)
			empRecord, err = ms.GetEmployeeByID(ctx, id)
			assert.Equal(t, err, nil)
			assert.Equal(t, empRecord.Position, oldPosition)
			oldSalary := empData.Salary

			updatedEmpData.Salary = 0.0
			err = ms.UpdateEmployee(ctx, id, updatedEmpData)
			assert.Equal(t, err, nil)
			empRecord, err = ms.GetEmployeeByID(ctx, id)
			assert.Equal(t, err, nil)
			assert.Equal(t, empRecord.Salary, oldSalary)

		})

		t.Run("Should return error when the id is not present", func(t *testing.T) {
			ms := NewMemoryStore()
			empData := &models.Employee{
				Name:     "Pavan",
				Position: "Developer",
				Salary:   11000,
			}
			id := 1111111
			err := ms.UpdateEmployee(ctx, id, empData)
			assert.Equal(t, err, errors.New(fmt.Sprintf("GetEmployeeByID: No employee found with id %d", id)))
		})
	})

	t.Run("Test Delete Employee By ID", func(t *testing.T) {
		t.Run("Should delete the employee record successfully when the id is present", func(t *testing.T) {
			ms := NewMemoryStore()
			empData := &models.Employee{
				Name:     "Pavan",
				Position: "Developer",
				Salary:   11000,
			}
			id, err := ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, nil)
			err = ms.DeleteEmployee(ctx, id)
			assert.Equal(t, err, nil)
			empRecord, err := ms.GetEmployeeByID(ctx, id)
			assert.Equal(t, err, errors.New(fmt.Sprintf("GetEmployeeByID: No employee found with id %d", id)))
			assert.Equal(t, empRecord, nil)
		})
		t.Run("Should return error when the id is not present", func(t *testing.T) {
			ms := NewMemoryStore()
			id := 111111
			err := ms.DeleteEmployee(ctx, id)
			assert.Equal(t, err, errors.New(fmt.Sprintf("GetEmployeeByID: No employee found with id %d", id)))
		})
	})

	t.Run("Test Get Employees", func(t *testing.T) {
		t.Run("Should get 1 employee records successfully when page is 1 and count is 1 for 2 records", func(t *testing.T) {
			ms := NewMemoryStore()
			empData := &models.Employee{
				Name:     "Pavan",
				Position: "Developer",
				Salary:   11000,
			}
			id, err := ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, nil)
			id, err = ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, nil)
			count := 1
			page := 1
			empData.ID = id
			empRecords, err := ms.GetEmployees(ctx, count, page)
			assert.Equal(t, err, nil)
			assert.Equal(t, empRecords[0], empData)
		})

		t.Run("Should get 0 employee records successfully when page is 3 and count is 1 for 2 records", func(t *testing.T) {
			ms := NewMemoryStore()
			empData := &models.Employee{
				Name:     "Pavan",
				Position: "Developer",
				Salary:   11000,
			}
			id, err := ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, nil)
			id, err = ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, nil)
			count := 1
			page := 3
			empData.ID = id
			empRecords, err := ms.GetEmployees(ctx, count, page)
			assert.Equal(t, err, nil)
			assert.Equal(t, len(empRecords), 0)
		})

		t.Run("Should get 2 employee records successfully when page is 1 and count is 3 for 2 records", func(t *testing.T) {
			ms := NewMemoryStore()
			empData := &models.Employee{
				Name:     "Pavan",
				Position: "Developer",
				Salary:   11000,
			}
			id, err := ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, nil)
			id, err = ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, nil)
			count := 3
			page := 1
			empData.ID = id
			empRecords, err := ms.GetEmployees(ctx, count, page)
			assert.Equal(t, err, nil)
			assert.Equal(t, len(empRecords), 2)
		})

		t.Run("Should get 2 employee records successfully when page is 1 and count is 3 for 2 records", func(t *testing.T) {
			ms := NewMemoryStore()
			empData := &models.Employee{
				Name:     "Pavan",
				Position: "Developer",
				Salary:   11000,
			}
			id, err := ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, nil)
			id, err = ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, nil)
			count := 3
			page := 1
			empData.ID = id
			empRecords, err := ms.GetEmployees(ctx, count, page)
			assert.Equal(t, err, nil)
			assert.Equal(t, len(empRecords), 2)
		})
	})

	t.Run("Test Get Total Employee Count", func(t *testing.T) {
		t.Run("Should get count of 2 for 2 records", func(t *testing.T) {
			ms := NewMemoryStore()
			empData := &models.Employee{
				Name:     "Pavan",
				Position: "Developer",
				Salary:   11000,
			}
			_, err := ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, nil)
			_, err = ms.CreateEmployee(ctx, empData)
			assert.Equal(t, err, nil)
			count, err := ms.GetTotalEmployeeCount(ctx)
			assert.Equal(t, err, nil)
			assert.Equal(t, count, 2)
		})
	})

}
