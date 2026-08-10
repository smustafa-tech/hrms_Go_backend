package repository

import (
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(emp *models.Employee) error {
	return r.db.Create(emp).Error
}

func (r *EmployeeRepository) FindAll() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.db.Where("deleted_at IS NULL").Find(&employees).Error
	return employees, err
}

func (r *EmployeeRepository) FindByEmpID(empID string) (*models.Employee, error) {
	var emp models.Employee
	err := r.db.Where("emp_id = ? AND deleted_at IS NULL", empID).First(&emp).Error
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *EmployeeRepository) Update(empID string, updates map[string]interface{}) (*models.Employee, error) {
	if err := r.db.Model(&models.Employee{}).
		Where("emp_id = ? AND deleted_at IS NULL", empID).
		Updates(updates).Error; err != nil {
		return nil, err
	}
	return r.FindByEmpID(empID)
}

func (r *EmployeeRepository) SoftDelete(empID string) error {
	return r.db.Model(&models.Employee{}).
		Where("emp_id = ?", empID).
		Update("deleted_at", gorm.Expr("NOW()")).Error
}

func (r *EmployeeRepository) EmpIDExists(empID string) bool {
	var count int64
	r.db.Model(&models.Employee{}).Where("emp_id = ?", empID).Count(&count)
	return count > 0
}

func (r *EmployeeRepository) EmailExists(email string) bool {
	var count int64
	r.db.Model(&models.Employee{}).Where("email = ? AND deleted_at IS NULL", email).Count(&count)
	return count > 0
}

func (r *EmployeeRepository) DepartmentCounts() []map[string]interface{} {
	type result struct {
		Department string
		Count      int64
	}
	var results []result
	r.db.Model(&models.Employee{}).
		Select("department, count(*) as count").
		Where("deleted_at IS NULL").
		Group("department").
		Scan(&results)

	counts := make([]map[string]interface{}, len(results))
	for i, r := range results {
		counts[i] = map[string]interface{}{
			"name":  r.Department,
			"value": r.Count,
		}
	}
	return counts
}
