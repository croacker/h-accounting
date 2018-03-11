package main

import (
	"fmt"
	"log"
	"path/filepath"

	"./commonutils"
	"./conf"
	"./filewatcher"
	"./ofd"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Загруженные чеки
var checks map[string]*ofd.OfdCheck

func printer(c chan string) {
	for {
		fullPath := <-c
		fileName := filepath.Base(fullPath)
		fmt.Println("Incoming file", fileName)
		ofdCheck, err := ofd.ReadCheck(fullPath)
		if err == nil {
			fmt.Println("Date time:", commonutils.ParseTimestamp(int64(ofdCheck.DateTime)))
			for idx, item := range ofdCheck.Items {
				fmt.Println("Item", idx, ":", item)
			}
			storeCheck(ofdCheck)
		} else {
			handleError(err)
		}
	}
}

func main() {
	appConf := conf.Get()
	fmt.Println("IncomingCheckFolder", appConf.IncomingCheckFolder)

	initPersistStore()

	var c = make(chan string)
	go printer(c)
	defer filewatcher.Watch("/home/alex/tmp/check", c).Close()
	doWait()
}

//Давай, жди
func doWait() {
	done := make(chan bool)
	<-done
}

//Инициализировать хранилище
func initPersistStore() {
	checks = make(map[string]*ofd.OfdCheck)
}

//Сохранить чек
func storeCheck(check *ofd.OfdCheck) {
	idx := string(check.DateTime)
	if checks[idx] == nil {
		checks[idx] = check
	}
	storeToMongo(check)
}

//Тестовый вариант записи в БД
func storeToMongo(check *ofd.OfdCheck) {
	config := conf.Get()
	dialInfo := mgo.DialInfo{
		Addrs:    []string{config.DbServer},
		Database: config.DbName,
		Username: config.DbUser,
		Password: config.DbPassword,
	}
	session, err := mgo.DialWithInfo(&dialInfo)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("accounting").C("testCheck")
	err = c.Insert(check)
	if err != nil {
		log.Fatal(err)
	}

	result := ofd.OfdCheck{}
	err = c.Find(bson.M{"user": "ООО \"О'КЕЙ\""}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Items from DB:", result.Items)
}

//Обработать ошибку
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
