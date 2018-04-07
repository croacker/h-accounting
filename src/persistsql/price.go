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
	Shop      Shop `gorm:"foreignkey:ShopId"`
	ShopId    uint
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
	dao.db.FirstOrCreate(price, Price{ShopId: price.Shop.ID, ProductId: price.Product.ID, DateTime: price.DateTime})
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
	dao.db.Where("product_id = ? AND shop_id = ? AND date_time = ?", price.ProductId, price.ShopId, price.DateTime).First(&dbPrice)
	// dao.db.Where(Price{ProductId: price.ProductId, ShopId:price.ShopId, DateTime:price.DateTime}).First(&dbPrice)
	return &dbPrice
}
