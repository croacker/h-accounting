package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"./conf"
	"./filewatcher"
	"./httpserver"
	"./ofd"
	// "./persistmongo"
	"./persistsql"
)

func printer(c chan string) {
	for {
		fullPath := <-c
		time.Sleep(1 * time.Second)
		fileName := filepath.Base(fullPath)
		fmt.Println("Incoming file", fileName)
		ofdChecks, err := ofd.ReadChecks(fullPath)
		if err == nil {
			store(ofdChecks)
		} else {
			handleError(err)
		}
	}
}

func main() {
	// persistsql.Connect()
	appConf := conf.Get()
	fmt.Println("IncomingCheckFolder", appConf.IncomingCheckFolder)

	var c = make(chan string)
	go printer(c)
	go httpserver.StartGin()
	defer filewatcher.Watch(appConf.IncomingCheckFolder, c).Close()
	doWait()
}

//Давай, жди
func doWait() {
	done := make(chan bool)
	<-done
}

//Тестовый вариант записи в БД
func store(check *ofd.OfdChecks) {
	fmt.Println("BEGIN save OFD checks")
	persistsql.Save(check)
	// persistmongo.Save(check)
	fmt.Println("Save OFD checks SUCCESS")
}

//Обработать ошибку
func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
