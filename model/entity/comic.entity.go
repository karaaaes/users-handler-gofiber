package entity

import (
	"time"

	"gorm.io/gorm"
)

type Comic struct {
	ID         int            `json:"id" gorm:"primaryKey"`
	SeriesName string         `json:"series_name" gorm:"type:varchar(255)"`
	Author     string         `json:"author" gorm:"type:varchar(255)"`
	Cover      string         `json:"cover"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
