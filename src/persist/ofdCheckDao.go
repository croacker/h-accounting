package persist

import (
	"../ofd"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DAO оригинального чека ОФД
type OfdCheckDao struct {
}

// Сохранить в БД
func (dao OfdCheckDao) Save(check *ofd.OfdCheck, session *mgo.Session) error {
	collection := getCollection("originalCheck", session)
	collection.Insert(check)
	return nil
}

//Найти в БД
func (dao OfdCheckDao) Find(check *ofd.OfdCheck, session *mgo.Session) *ofd.OfdCheck {
	collection := getCollection("originalCheck", session)
	result := ofd.OfdCheck{}
	collection.Find(bson.M{"datetime": check.DateTime, "userinn": check.UserInn}).One(&result)
	return &result
}

//Удалить из БД
func (dao OfdCheckDao) Delete(check *ofd.OfdCheck, session *mgo.Session) error {
	collection := getCollection("originalCheck", session)
	collection.Remove(bson.M{"datetime": check.DateTime, "userinn": check.UserInn})
	return nil
}

//Получить
func (dao OfdCheckDao) GetAll(session *mgo.Session) error {
	collection := getCollection("originalCheck", session)
	var results []ofd.OfdCheck
	return collection.Find(nil).All(results)
}
