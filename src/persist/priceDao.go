package persist

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DAO сущности цена на товар у указанного продавца в указанный момент времени
type PriceDao struct {
}

// Сохранить в БД
func (dao PriceDao) Save(price *Price, session *mgo.Session) (*Price, error) {
	resultPrice, err := dao.Find(price, session)
	if err != nil {
		price.Id = bson.NewObjectId()
		collection := collection("price", session)
		collection.Insert(price)
		resultPrice = price
	}
	return resultPrice, nil
}

//Найти в БД
func (dao PriceDao) Find(price *Price, session *mgo.Session) (*Price, error) {
	collection := collection("price", session)
	result := Price{}
	err := collection.Find(bson.M{"goodsid": price.GoodsId, "shopid": price.ShopId,
		"datetime": price.DateTime}).One(&result)
	return &result, err
}

//Найти в БД
func (dao PriceDao) FindId(price *Price, session *mgo.Session) (*Price, error) {
	collection := collection("Price", session)
	result := Price{}
	err := collection.FindId(price.Id).One(&result)
	return &result, err
}

//Удалить из БД
func (dao PriceDao) Delete(price *Price, session *mgo.Session) error {
	collection := collection("Price", session)
	collection.Remove(bson.M{"goodsid": price.GoodsId, "shopid": price.ShopId,
		"datetime": price.DateTime})
	return nil
}

//Получить
func (dao PriceDao) GetAll(session *mgo.Session) ([]Price, error) {
	collection := collection("price", session)
	var results []Price
	err := collection.Find(nil).All(results)
	return results, err
}
