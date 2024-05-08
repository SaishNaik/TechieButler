package models

type GetEmployeeRecordsDTO struct {
	Records []*Employee
	Error   error
}

type GetEmployeeCountDTO struct {
	Count int
	Error error
}
