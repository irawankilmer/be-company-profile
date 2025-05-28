package domain

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100"`
	Username  string `gorm:"uniqueIndex;size:50"`
	Email     string `gorm:"uniqueIndex;size:100"`
	Password  string `gorm:"size:255"`
	Roles     []Role `gorm:"many2many:user_roles"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
