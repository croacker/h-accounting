package persistsql

import "github.com/jinzhu/gorm"

type ProductCathegory struct {
	gorm.Model
	Name     string
	Products []Product `gorm:"foreignkey:ProductCathegoryId"`
}
