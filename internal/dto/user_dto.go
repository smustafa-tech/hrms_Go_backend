package dto

type UpdateProfileRequest struct {
	FirstName   string `json:"firstName"`
	MiddleName  string `json:"middleName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	AdharCard   string `json:"adharCard"`
	Designation string `json:"designation"`
	Bio         string `json:"bio"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=6"`
}
