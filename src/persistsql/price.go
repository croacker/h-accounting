package persistsql

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Цена на товар у указанного продавца
type Price struct {
	gorm.Model
	Product   Product `gorm:"foreignkey:ProductId"`
	ProductId uint
	Sailer    Sailer `gorm:"foreignkey:SailerId"`
	SailerId  uint
	Price     int
	DateTime  int
}

type PriceDao struct {
	db *gorm.DB
}

func (dao PriceDao) Create(price *Price) {
	dao.db.Create(price)
}

func (dao PriceDao) FirstOrCreate(price *Price) *Price {
	dao.db.FirstOrCreate(price, Price{SailerId: price.Sailer.ID, ProductId: price.Product.ID, DateTime: price.DateTime})
	fmt.Println("Price Id:", price.ID)
	return price
}

func (dao PriceDao) Save(price *Price) {
	dao.db.Save(price)
}

func (dao PriceDao) FindById(id uint) *Price {
	var price Price
	dao.db.First(&price, id)
	return &price
}

func (dao PriceDao) Find(price *Price) *Price {
	var dbPrice Price
	dao.db.Where("product_id = ? AND sailer_id = ? AND date_time = ?", price.ProductId, price.SailerId, price.DateTime).First(&dbPrice)
	// dao.db.Where(Price{ProductId: price.ProductId, ShopId:price.ShopId, DateTime:price.DateTime}).First(&dbPrice)
	return &dbPrice
}
