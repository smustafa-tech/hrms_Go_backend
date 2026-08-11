package dto

type ApplyLeaveRequest struct {
	LeaveType string `json:"leaveType" binding:"required"`
	StartDate string `json:"startDate" binding:"required"`
	EndDate   string `json:"endDate" binding:"required"`
	TotalDays int    `json:"totalDays"`
	Reason    string `json:"reason"`
}

type UpdateLeaveRequest struct {
	LeaveID string `json:"leaveId" binding:"required"`
	Status  string `json:"status" binding:"required"`
}
