package models

import (
	"time"

	"github.com/CavasCallahan/firstGo/server/services"
	"gorm.io/gorm"
)

type Base struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (base *Base) BeforeCreate(thx *gorm.DB) (err error) {
	base.ID = services.GenerateUuidv4()
	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
	return
}
