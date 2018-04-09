package persistsql

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Тип магазин
type Shop struct {
	gorm.Model
	Sailer   Sailer `gorm:"foreignkey:SailerId"`
	SailerId uint
	//Наименование
	Name    string
	Address string
}

type ShopDao struct {
	db *gorm.DB
}

func (dao ShopDao) Create(shop *Shop) {
	dao.db.Create(shop)
}

func (dao ShopDao) FirstOrCreate(sailer *Sailer, address string) *Shop {
	shop := &Shop{Sailer: *sailer,
		Name:    toName(sailer, address),
		Address: address}
	dao.db.FirstOrCreate(shop, Shop{SailerId: sailer.ID, Address: address})
	fmt.Println("Shop Id:", sailer.ID)
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

func (dao ShopDao) GetAll() []Shop {
	var shop []Shop
	dao.db.Find(&shop)
	return shop
}

func toName(sailer *Sailer, address string) string {
	result := sailer.Name
	if len(address) != 0 {
		result += "(" + address + ")"
	}
	return result
}
