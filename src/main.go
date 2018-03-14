package main

import (
	"fmt"
	"log"
	"path/filepath"

	"./commonutils"
	"./conf"
	"./filewatcher"
	"./ofd"
	"./persist"
)

func printer(c chan string) {
	for {
		fullPath := <-c
		fileName := filepath.Base(fullPath)
		fmt.Println("Incoming file", fileName)
		ofdCheck, err := ofd.ReadCheck(fullPath)
		if err == nil {
			fmt.Println("Date time:", commonutils.ParseTimestamp(int64(ofdCheck.DateTime)))
			// for idx, item := range ofdCheck.Items {
			// 	fmt.Println("Item", idx, ":", item)
			// }
			storeToMongo(ofdCheck)
		} else {
			handleError(err)
		}
	}
}

func main() {
	appConf := conf.Get()
	fmt.Println("IncomingCheckFolder", appConf.IncomingCheckFolder)

	var c = make(chan string)
	go printer(c)
	defer filewatcher.Watch(appConf.IncomingCheckFolder, c).Close()
	doWait()
}

//Давай, жди
func doWait() {
	done := make(chan bool)
	<-done
}

//Тестовый вариант записи в БД
func storeToMongo(check *ofd.OfdCheck) {
	persist.Save(check)
}

//Обработать ошибку
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
