package dto

type CreateEmployeeRequest struct {
	EmpID         string   `json:"emp_id" binding:"required"`
	FirstName     string   `json:"firstName" binding:"required"`
	MiddleName    string   `json:"middleName"`
	LastName      string   `json:"lastName" binding:"required"`
	Email         string   `json:"email" binding:"required,email"`
	Phone         string   `json:"phone"`
	AdharCard     string   `json:"adharCard"`
	Designation   string   `json:"designation"`
	Role          string   `json:"role"`
	Department    string   `json:"department"`
	DateOfJoining string   `json:"dateOfJoining"`
	Status        string   `json:"status"`
	Salary        *float64 `json:"salary"`
	WorkMode      string   `json:"workMode"`
	MgrID         string   `json:"mgrId"`
	HrID          string   `json:"hrId"`
}

type UpdateEmployeeRequest struct {
	FirstName     string   `json:"firstName"`
	MiddleName    string   `json:"middleName"`
	LastName      string   `json:"lastName"`
	Email         string   `json:"email"`
	Phone         string   `json:"phone"`
	AdharCard     string   `json:"adharCard"`
	Designation   string   `json:"designation"`
	Role          string   `json:"role"`
	Department    string   `json:"department"`
	DateOfJoining string   `json:"dateOfJoining"`
	Status        string   `json:"status"`
	Salary        *float64 `json:"salary"`
	WorkMode      string   `json:"workMode"`
	MgrID         string   `json:"mgrId"`
	HrID          string   `json:"hrId"`
}
