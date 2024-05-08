package repo

import (
	"context"
	"techiebutler/models"
)

//go:generate go run -mod=mod github.com/golang/mock/mockgen  -source=repo.go --build_flags=--mod=mod -destination=./mock_repo.go -package=repo . IRepo
type IRepo interface {
	CreateEmployee(ctx context.Context, emp *models.Employee) (int, error)
	GetEmployeeByID(ctx context.Context, id int) (emp *models.Employee, err error)
	GetEmployees(ctx context.Context, count, page int) ([]*models.Employee, error)
	GetTotalEmployeeCount(ctx context.Context) (int, error)
	UpdateEmployee(ctx context.Context, id int, emp *models.Employee) error
	DeleteEmployee(ctx context.Context, id int) error
}
