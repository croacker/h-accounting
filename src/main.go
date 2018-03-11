package main

import (
	"fmt"
	"log"
	"path/filepath"

	"./commonutils"
	"./conf"
	"./filewatcher"
	"./ofd"
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
}

//Обработать ошибку
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
