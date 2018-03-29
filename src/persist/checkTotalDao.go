package persist

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DAO сущности товар
type CheckTotalDao struct {
}

// Сохранить в БД
func (dao CheckTotalDao) Save(checkTotal *CheckTotal, session *mgo.Session) (*CheckTotal, error) {
	resultCheckTotal, err := dao.FindName(checkTotal, session)
	if err != nil {
		checkTotal.Id = bson.NewObjectId()
		collection := collection("checkTotal", session)
		collection.Insert(checkTotal)
		resultCheckTotal = checkTotal
	}
	return resultCheckTotal, nil
}

//Найти в БД
func (dao CheckTotalDao) FindName(checkTotal *CheckTotal, session *mgo.Session) (*CheckTotal, error) {
	collection := collection("checkTotal", session)
	result := CheckTotal{}
	err := collection.Find(bson.M{"shopid": checkTotal.ShopId, "datetime": checkTotal.DateTime}).One(&result)
	return &result, err
}

//Найти в БД
func (dao CheckTotalDao) FindId(checkTotal *CheckTotal, session *mgo.Session) (*CheckTotal, error) {
	collection := collection("checkTotal", session)
	result := CheckTotal{}
	err := collection.FindId(checkTotal.Id).One(&result)
	return &result, err
}

//Удалить из БД
func (dao CheckTotalDao) Delete(checkTotal *CheckTotal, session *mgo.Session) error {
	collection := collection("checkTotal", session)
	collection.Remove(bson.M{"shopid": checkTotal.ShopId, "datetime": checkTotal.DateTime})
	return nil
}

//Получить
func (dao CheckTotalDao) GetAll(session *mgo.Session) ([]CheckTotal, error) {
	collection := collection("checkTotal", session)
	var results []CheckTotal
	err := collection.Find(nil).All(&results)
	return results, err
}
