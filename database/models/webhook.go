package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Webhook struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt pq.NullTime
}
