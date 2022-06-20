package requests

import "gorm.io/datatypes"

type Lead struct {
	Name        string          `binding:"required"`
	Email       string          `binding:"required"`
	Phone       string          `binding:"required"`
	OtherFields *datatypes.JSON `json:"other_fields"`
}
