package domain

import "time"

type Post struct {
	ID         uint   `gorm:"primaryKey"`
	Title      string `gorm:"size:255"`
	Content    string `gorm:"type:text"`
	UserID     uint
	Categories []Category `gorm:"many2many:category_posts"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
