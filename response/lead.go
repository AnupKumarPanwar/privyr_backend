package response

import (
	"time"

	"gorm.io/datatypes"
)

type Lead struct {
	Name        string          `json:"name"`
	Email       string          `json:"email"`
	Phone       string          `json:"phone"`
	OtherFields *datatypes.JSON `json:"other_fields"`
	CreatedAt   time.Time       `json:"created_at"`
}
