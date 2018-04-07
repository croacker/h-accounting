package persistsql

import (
	"fmt"

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
	db *gorm.DB
}

func (dao ShopDao) Create(shop *Shop) {
	dao.db.Create(shop)
}

func (dao ShopDao) CreateIfNotExists(name string, inn string) *Shop {
	shop := dao.FindByInn(inn)
	if shop.ID == 0 {
		shop = &Shop{Name: name, Inn: inn}
		dao.Create(shop)
		fmt.Println("New shop Id:", shop.ID)
	} else {
		fmt.Println("Shop exists Id:", shop.ID)
	}
	return shop
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
