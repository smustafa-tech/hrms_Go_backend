package dto

type PayrollRequest struct {
	UserID      string  `json:"userId" binding:"required"`
	Month       int     `json:"month" binding:"required,min=1,max=12"`
	Year        int     `json:"year" binding:"required"`
	BasicSalary float64 `json:"basicSalary" binding:"required"`
	HRA         float64 `json:"hra"`
	Allowances  float64 `json:"allowances"`
	Deductions  float64 `json:"deductions"`
}

type PayrollStatusRequest struct {
	Status string `json:"status" binding:"required"`
}
