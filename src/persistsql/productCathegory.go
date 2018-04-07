package persistsql

import "github.com/jinzhu/gorm"

type ProductCathegory struct {
	gorm.Model
	Name     string
	Products []Product `gorm:"foreignkey:ProductCathegoryId"`
}

type ProductCathegoryDao struct {
	db gorm.DB
}

func (dao ProductCathegoryDao) Create(cathegory *ProductCathegory) {
	dao.db.Create(cathegory)
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
