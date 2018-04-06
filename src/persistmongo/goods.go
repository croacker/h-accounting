package persistmongo

import (
	"gopkg.in/mgo.v2/bson"
)

//Тип товар, названия полей сохранены как в оригинальном json
type Goods struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string
}

//Фабричный метод
func NewGoods(name string) *Goods {
	return &Goods{Name: name}
}
