package main

import (
	"fmt"
	"log"

	"./conf"
	"./emailprocess"
	"./fileprocess"
	"./filewatcher"
	"./httpserver"
	"./ofd"
	"./persistsql"
)

func main() {
	appConf := conf.Get()
	emailprocess.Receive()
	return

	persistsql.Init()
	fmt.Println("IncomingCheckFolder", appConf.IncomingCheckFolder)

	var fileNameChanel = make(chan string)

	go httpserver.StartGin()
	go fileprocess.ProcessFile(fileNameChanel, store)

	defer filewatcher.Watch(appConf.IncomingCheckFolder, fileNameChanel).Close()
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
