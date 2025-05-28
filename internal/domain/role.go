package domain

type Role struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"uniqueIndex;size:50"`
	Users []User `gorm:"many2many:user_roles"`
}
