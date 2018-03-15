package persist

import (
	"fmt"
	"log"

	"../conf"
	"../ofd"
	"gopkg.in/mgo.v2"
)

// Save - Сохранить чек в хранилище
func Save(check *ofd.OfdCheck) error {
	var err error
	session, err := getSession()
	handleError(err)
	if err == nil {
		defer session.Close()
		//Сохранить оригинальный чек
		ofdCheckDao := OfdCheckDao{}
		ofdCheckDao.Save(check, session)

		shop := NewShop(check)
		shopDao := ShopDao{}
		shop, _ = shopDao.Save(shop, session)
		fmt.Println("Save Shop id:", shop.Id.String(), ", user:", shop.User, " INN:", shop.UserInn)

		goodsDao := GoodsDao{}
		priceDao := PriceDao{}
		for _, item := range check.Items {
			goods := NewGoods(item.Name)
			goods, _ = goodsDao.Save(goods, session)
			fmt.Println("Save Goods id:", goods.Id.String(), ", name:", goods.Name)

			price := NewPrice(goods, shop, item.Price, check.DateTime)
			price, _ = priceDao.Save(price, session)
			fmt.Println("Save Price id:", price.Id.String(), ", Good name:", goods.Name)
		}
		checkTotalDao := CheckTotalDao{}
		checkTotal := NewCheckTotal(shop, check)
		checkTotal, _ = checkTotalDao.Save(checkTotal, session)
		fmt.Println("Save Check total id:", checkTotal.Id.String())
	}
	return err
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
func collection(name string, session *mgo.Session) *mgo.Collection {
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
