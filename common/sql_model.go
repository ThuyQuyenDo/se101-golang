package common

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseEntity struct {
	ID     *string `json:"id" gorm:"column:id;"`
	Status string  `json:"status" gorm:"column:status;default:active;"`

	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

// Hook
func (record *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {
	if record.ID == nil {
		id := uuid.New().String()
		record.ID = &id
	}

	return
}
