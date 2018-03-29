package persist

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DAO сущности товар
type GoodsDao struct {
}

// Сохранить в БД
func (dao GoodsDao) Save(goods *Goods, session *mgo.Session) (*Goods, error) {
	resultGoods, err := dao.FindName(goods, session)
	if err != nil {
		goods.Id = bson.NewObjectId()
		collection := collection("goods", session)
		collection.Insert(goods)
		resultGoods = goods
	}
	return resultGoods, nil
}

//Найти в БД
func (dao GoodsDao) FindName(goods *Goods, session *mgo.Session) (*Goods, error) {
	collection := collection("goods", session)
	result := Goods{}
	err := collection.Find(bson.M{"name": goods.Name}).One(&result)
	return &result, err
}

//Найти в БД
func (dao GoodsDao) FindId(goods *Goods, session *mgo.Session) (*Goods, error) {
	collection := collection("goods", session)
	result := Goods{}
	err := collection.FindId(goods.Id).One(&result)
	return &result, err
}

//Удалить из БД
func (dao GoodsDao) Delete(goods *Goods, session *mgo.Session) error {
	collection := collection("goods", session)
	collection.Remove(bson.M{"name": goods.Name})
	return nil
}

//GetAll Получить
func (dao GoodsDao) GetAll(session *mgo.Session) ([]Goods, error) {
	collection := collection("goods", session)
	var results []Goods
	err := collection.Find(nil).All(&results)
	return results, err
}
