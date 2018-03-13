package persist

import (
	"fmt"
	"log"

	"../conf"
	"../ofd"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Save - Сохранить чек в хранилище
func Save(check *ofd.OfdCheck) error {
	var err error
	session, err := getSession()
	handleError(err)
	if err == nil {
		defer session.Close()
		//Сохранить оригинальный чек
		dao := OfdCheckDao{}
		dao.Save(check, session)

		shop := NewShop(check)
		dao = ShopDao{}
		err = persistShop(shop, session)
		handleError(err)

		findedShop, err := findShopByID(shop.Id, session)
		handleError(err)
		if err == nil {
			fmt.Println("findedShop", findedShop)
		}
	}
	return err
}

//Сохранить в БД информацию о продавце(магазине)
func persistShop(shop *ofd.Shop, session *mgo.Session) error {
	shop.Id = bson.NewObjectId()
	collection := session.DB(conf.Get().DbName).C("shop")
	return collection.Insert(shop)
}

// Найти в БД информацию о продавце(магазине) по id
func findShopByID(id bson.ObjectId, session *mgo.Session) (*ofd.Shop, error) {
	shop := ofd.Shop{}
	collection := session.DB(conf.Get().DbName).C("shop")
	err := collection.FindId(id).One(&shop)
	return &shop, err
}

//Получить сессию
func getSession() (*mgo.Session, error) {
	dialInfo := getDialInfo()
	session, err := mgo.DialWithInfo(dialInfo)
	return session, err
}

//Получить БД
func getDatabase(session *mgo.Session) *mgo.Database {
	return session.DB(conf.Get().DbName)
}

//Получить коллекцию
func getCollection(name string, session *mgo.Session) *mgo.Collection {
	db := getDatabase(session)
	return db.C(name)
}

//Получить данные для доступа к хранилищу
func getDialInfo() *mgo.DialInfo {
	config := conf.Get()
	dialInfo := mgo.DialInfo{
		Addrs:    []string{config.DbServer},
		Database: config.DbName,
		Username: config.DbUser,
		Password: config.DbPassword,
	}
	return &dialInfo
}

//Обработать ошибку
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
