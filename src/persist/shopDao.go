package persist

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DAO сущности магазни
type ShopDao struct {
}

// Сохранить в БД
func (dao ShopDao) Save(shop *Shop, session *mgo.Session) error {
	_, err := dao.Find(shop, session)
	if err != nil {
		shop.Id = bson.NewObjectId()
		collection := collection("shop", session)
		collection.Insert(shop)
	}
	return nil
}

//Найти в БД
func (dao ShopDao) Find(shop *Shop, session *mgo.Session) (*Shop, error) {
	collection := collection("shop", session)
	result := Shop{}
	err := collection.Find(bson.M{"userinn": shop.UserInn}).One(&result)
	return &result, err
}

//Удалить из БД
func (dao ShopDao) Delete(shop *Shop, session *mgo.Session) error {
	collection := collection("shop", session)
	collection.Remove(bson.M{"userinn": shop.UserInn})
	return nil
}

//Получить
func (dao ShopDao) GetAll(session *mgo.Session) ([]Shop, error) {
	collection := collection("shop", session)
	var results []Shop
	err := collection.Find(nil).All(results)
	return results, err
}
