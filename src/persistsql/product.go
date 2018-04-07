package persistsql

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Cathegory   ProductCathegory `gorm:"foreignkey:CathegoryId"`
	CathegoryId uint
}

type ProductDao struct {
	db *gorm.DB
}

func (dao ProductDao) Create(product *Product) {
	dao.db.Create(product)
}

func (dao ProductDao) FirstOrCreate(name string, cathegory *ProductCathegory) *Product {
	var product = &Product{Name: name, Cathegory: *cathegory}
	dao.db.FirstOrCreate(product, Product{Name: name})
	fmt.Println("Product Id:", product.ID)
	return product
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

func (dao ProductDao) GetAll() []Product {
	var products []Product
	dao.db.Preload("Cathegory").Find(&products)
	return products
}
