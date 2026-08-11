package dto

type NotificationRequest struct {
	SenderID     string `json:"senderId" binding:"required"`
	ReceiverID   string `json:"receiverId" binding:"required"`
	SenderType   string `json:"senderType" binding:"required"`
	ReceiverType string `json:"receiverType" binding:"required"`
	Type         string `json:"type" binding:"required"`
	Message      string `json:"message" binding:"required"`
}
