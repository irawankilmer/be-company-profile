package domain

import "time"

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:255"`
	Content   string `gorm:"type:text"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
