package domain

import "time"

type Category struct {
	ID               uint   `gorm:"primaryKey"`
	Name             string `gorm:"size:125"`
	Description      string `gorm:"type:text"`
	ParentCategoryID *uint
	ParentCategory   *Category  `gorm:"foreignKey:ParentCategoryID"`
	SubCategories    []Category `gorm:"foreignKey:ParentCategoryID"`
	Posts            []Post     `gorm:"many2many:category_posts"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
