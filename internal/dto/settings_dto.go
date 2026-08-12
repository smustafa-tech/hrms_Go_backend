package dto

type NotificationSettingsRequest struct {
	EmailNotifications *bool `json:"emailNotifications"`
	PushNotifications  *bool `json:"pushNotifications"`
	LeaveRequests      *bool `json:"leaveRequests"`
	AttendanceAlerts   *bool `json:"attendanceAlerts"`
}

type SecuritySettingsRequest struct {
	TwoFactorAuth  *bool  `json:"twoFactorAuth"`
	SessionTimeout string `json:"sessionTimeout"`
}
