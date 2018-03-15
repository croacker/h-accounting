package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"./conf"
	"./filewatcher"
	"./ofd"
	"./persist"
)

func printer(c chan string) {
	for {
		fullPath := <-c
		time.Sleep(1 * time.Second)
		fileName := filepath.Base(fullPath)
		fmt.Println("Incoming file", fileName)
		ofdCheck, err := ofd.ReadCheck(fullPath)
		if err == nil {
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
	fmt.Println("BEGIN save OFD check")
	persist.Save(check)
	fmt.Println("Save OFD check SUCCESS")
}

//Обработать ошибку
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
