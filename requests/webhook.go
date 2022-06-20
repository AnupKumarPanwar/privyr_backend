package requests

import "github.com/google/uuid"

type Webhook struct {
	Name   string    `binding:"required"`
	UserID uuid.UUID `json:"user_id" binding:"required"`
}
