package persist

import (
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
		resultShop, _ := shopDao.Save(shop, session)

		goodsDao := GoodsDao{}
		priceDao := PriceDao{}
		for _, item := range check.Items {
			goods := NewGoods(item.Name)
			goods, _ = goodsDao.Save(goods, session)

			price := NewPrice(goods, resultShop, item.Price, check.DateTime)
			priceDao.Save(price, session)
		}
		checkTotal := NewCheckTotal(shop, check)
		checkTotalDao := CheckTotalDao{}
		checkTotalDao.Save(checkTotal, session)
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
