package main

import (
	"fmt"
	"log"

	"./conf"
	"./fileprocess"
	"./filewatcher"
	"./httpserver"
	"./ofd"
	"./persistsql"
)

func main() {
	persistsql.Init()
	appConf := conf.Get()
	fmt.Println("IncomingCheckFolder", appConf.IncomingCheckFolder)

	var c = make(chan string)

	go fileprocess.ProcessFile(c, store)
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
	fmt.Println("Save OFD checks SUCCESS")
}

//Обработать ошибку
func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
