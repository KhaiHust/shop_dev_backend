package model

import "time"

type BaseModel struct {
	ID        uint       `gorm:"column:id; primary_key"`
	CreatedAt *time.Time `gorm:"column:created_at; type: timestamp; auto_now_add"`
	UpdatedAt *time.Time `gorm:"column:updated_at; type: timestamp; auto_now"`
}
