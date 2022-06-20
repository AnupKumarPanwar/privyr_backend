package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt pq.NullTime
}
