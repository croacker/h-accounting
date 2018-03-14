package persist

import (
	"gopkg.in/mgo.v2/bson"
)

//Цена на товар у указанного продавца
type Price struct {
	Id       bson.ObjectId `bson:"_id"`
	GoodsId  bson.ObjectId
	ShopId   bson.ObjectId
	Price    int
	DateTime int
}


func NewPrice(goods Goods, shop Shop, price int, datatime int) 