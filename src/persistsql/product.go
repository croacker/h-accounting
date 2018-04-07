package persistsql

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name               string
	Cathegory          ProductCathegory `gorm:"foreignkey:ProductCathegoryId"`
	ProductCathegoryId uint
}

type ProductDao struct {
	db gorm.DB
}

func (dao ProductDao) Create(product *Product) {
	dao.db.Create(product)
}

func (dao ProductDao) Save(product *Product) {
	dao.db.Save(product)
}

func (dao ProductDao) FindById(id uint) *Product {
	var product Product
	dao.db.First(&product, id)
	return &product
}

func (dao ProductDao) FindByName(name string) *Product {
	var product Product
	dao.db.First(&product, "name = ?", name)
	return &product
}
