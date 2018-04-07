package persistsql

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type ProductCathegory struct {
	gorm.Model
	Name     string
	Products []Product `gorm:"foreignkey:ProductCathegoryId"`
}

type ProductCathegoryDao struct {
	db *gorm.DB
}

func (dao ProductCathegoryDao) Create(cathegory *ProductCathegory) {
	dao.db.Create(cathegory)
}

func (dao ProductCathegoryDao) CreateIfNotExists(name string) *ProductCathegory {
	productCathegory := dao.FindByName(name)
	if productCathegory.ID == 0 {
		productCathegory = &ProductCathegory{Name: name}
		dao.Create(productCathegory)
		fmt.Println("New product cathegory Id:", productCathegory.ID)
	} else {
		fmt.Println("Product cathegory exists Id:", productCathegory.ID)
	}
	return productCathegory
}

func (dao ProductCathegoryDao) Save(cathegory *ProductCathegory) {
	dao.db.Save(cathegory)
}

func (dao ProductCathegoryDao) FindById(id uint) *ProductCathegory {
	var cathegory ProductCathegory
	dao.db.First(&cathegory, id)
	return &cathegory
}

func (dao ProductCathegoryDao) FindByName(name string) *ProductCathegory {
	var cathegory ProductCathegory
	dao.db.First(&cathegory, "name = ?", name)
	return &cathegory
}
