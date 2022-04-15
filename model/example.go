package model

import (
	"time"

	"gorm.io/gorm"
)

// swagger:model ServiceTree
type Example struct {
	// swagger:ignore
	gorm.Model
	Path      string
	NodeName  string
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Example) TableName() string {
	return "service_tree"
}
