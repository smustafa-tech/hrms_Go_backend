package dto

type MarkAttendanceRequest struct {
	UserID string `json:"userId" binding:"required"`
	Date   string `json:"date" binding:"required"`
	Status string `json:"status"`
	Action string `json:"action"`
}

type AdminMarkAttendanceRequest struct {
	UserID string `json:"userId" binding:"required"`
	Date   string `json:"date" binding:"required"`
	Status string `json:"status"`
}

type BreakRequest struct {
	AttendanceID string `json:"attendanceId" binding:"required"`
}
