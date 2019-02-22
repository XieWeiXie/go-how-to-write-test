package partfour

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	Name string `gorm:"type:varchar" json:"name"`
}

type Items []Item

func NewItems() Items {
	return []Item{
		{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "XieWei",
		},
		{
			Model: gorm.Model{
				ID:        2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "WuXiaoShen",
		},
	}
}
