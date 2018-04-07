package persistsql

import (
	"github.com/jinzhu/gorm"
)

//Тип продавец(магазин)
type Shop struct {
	gorm.Model
	//Наименование
	Name string
	//ИНН
	Inn string
}

type ShopDao struct {
	db gorm.DB
}

func (dao ShopDao) Create(shop *Shop) {
	dao.db.Create(shop)
}

func (dao ShopDao) Save(shop *Shop) {
	dao.db.Save(shop)
}

func (dao ShopDao) FindById(id uint) *Shop {
	var shop Shop
	dao.db.First(&shop, id)
	return &shop
}

func (dao ShopDao) FindByInn(inn string) *Shop {
	var shop Shop
	dao.db.First(&shop, "inn = ?", inn)
	return &shop
}