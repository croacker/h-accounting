package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"./conf"
	"./filewatcher"
)

func printer(c chan string) {
	for {
		fullPath := <-c
		fileName := filepath.Base(fullPath)
		fmt.Println("File event", fileName)

	}
}

func main() {
	appConf := conf.Get()
	fmt.Println(appConf.IncomingCheckFolder)

	fmt.Println(parseTimestamp(1520328120))
	var c = make(chan string)
	go printer(c)
	defer filewatcher.Watch("/home/alex/tmp/check", c).Close()
	doWait()
}

func doWait() {
	done := make(chan bool)
	<-done
}

/**
*Преобразовать timestamp из чека(например ofdCheck.DateTime) во время.
 */
func parseTimestamp(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

/**
*Обработать ошибку
 */
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
