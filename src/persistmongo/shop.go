package persistmongo

import (
	"../ofd"
	"gopkg.in/mgo.v2/bson"
)

//Тип продавец(магазин), названия атрибутов сохранены как в оригинальном json
type Shop struct {
	Id bson.ObjectId `bson:"_id"`
	//Наименование
	User string
	//ИНН
	UserInn string
}

//Фабричный метод
func NewShop(check *ofd.OfdCheck) *Shop {
	return &Shop{User: check.User, UserInn: check.UserInn}
}
