package model

import "gorm.io/gorm"

type Novel struct {
	ID          int    `gorm:"type:int;primary_key" json:"id"`
	Name        string `gorm:"type:varchar(50)" json:"name"`
	Author      string `gorm:"type:varchar(50)" json:"author"`
	Description string `gorm:"type:varchar(50)" json:"description"`
	*gorm.Model
}
