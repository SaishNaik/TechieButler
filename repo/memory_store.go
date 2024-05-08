package repo

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"techiebutler/models"
)

// MemoryStore using a map to facilitate deletion

type MemoryStore struct {
	records map[int]*models.Employee
	counter int64
	mutex   sync.Mutex
}

func (m *MemoryStore) GetTotalEmployeeCount(ctx context.Context) (int, error) {
	//TODO implement me
	return len(m.records), nil
}

func (m *MemoryStore) GetEmployees(ctx context.Context, count, page int) ([]*models.Employee, error) {
	keys := make([]int, 0)
	for k, _ := range m.records {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	start := (page - 1) * count
	end := (page-1)*count + count
	empRecords := make([]*models.Employee, 0)
	if end > len(keys) {
		end = len(keys)
	}
	if start < len(keys) {
		for _, k := range keys[start:end] {
			empRecords = append(empRecords, m.records[k])
		}
	}

	return empRecords, nil
}

func (m *MemoryStore) getId() int {
	return int(atomic.AddInt64(&m.counter, 1))
}

func (m *MemoryStore) CreateEmployee(ctx context.Context, emp *models.Employee) (int, error) {
	if strings.TrimSpace(emp.Name) == "" || emp.Salary <= 0.0 || strings.TrimSpace(emp.Position) == "" {
		return 0, errors.New("name,salary or position is empty or incorrect")
	}
	id := m.getId()
	emp.ID = id
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.records[id] = emp
	return id, nil
}

func (m *MemoryStore) GetEmployeeByID(ctx context.Context, id int) (emp *models.Employee, err error) {
	empRecord, ok := m.records[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("GetEmployeeByID: No employee found with id %d", id))
	}
	return empRecord, nil
}

func (m *MemoryStore) UpdateEmployee(ctx context.Context, id int, newData *models.Employee) error {
	empRecord, err := m.GetEmployeeByID(ctx, id)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if empRecord != nil {
		if strings.TrimSpace(newData.Name) != "" {
			m.records[id].Name = newData.Name
		}
		if strings.TrimSpace(newData.Position) != "" {
			m.records[id].Position = newData.Position
		}
		if newData.Salary > 0.0 {
			m.records[id].Salary = newData.Salary
		}
	}
	return err
}

func (m *MemoryStore) DeleteEmployee(ctx context.Context, id int) error {
	empRecord, err := m.GetEmployeeByID(ctx, id)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if empRecord != nil {
		delete(m.records, id)
	}
	return err
}

//func (m *MemoryStore) employeeExists(id int) (*models.Employee,error){
//	empRecord, ok := m.records[id]
//	if !ok {
//		return nil,errors.New(fmt.Sprintf("GetEmployeeByID: No employee found with id %d", id))
//	}
//	return empRecord,nil
//}

func NewMemoryStore() IRepo {
	return &MemoryStore{
		records: make(map[int]*models.Employee),
		mutex:   sync.Mutex{},
	}
}
