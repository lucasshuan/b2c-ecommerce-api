package model

import (
	"encoding/json"
	"io"
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        uint            `json:"id" gorm:"primary_key"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

func (b *Base) UnmarshalGQL() error {
	return nil
}
func (b Base) MarshalGQL(w io.Writer) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":         b.ID,
		"created_at": b.CreatedAt,
		"updated_at": b.UpdatedAt,
	})
}
