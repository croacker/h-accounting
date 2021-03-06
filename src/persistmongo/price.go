package persistmongo

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

//Фабричный метод
func NewPrice(goods *Goods, shop *Shop, checkPrice int, datatime int) *Price {
	return &Price{bson.NewObjectId(), goods.Id, shop.Id, checkPrice, datatime}
}
