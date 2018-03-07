package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"./filewatcher"
	"./ofd"
)

func printer(c chan string) {
	for {
		msg := <-c
		file := filepath.Base(msg)
		fmt.Println("File event", file)
	}
}

func main() {
	fmt.Println(parseTimestamp(1520328120))
	fmt.Println(parseTimestamp(106374515))
	var c = make(chan string)
	go printer(c)
	defer filewatcher.Watch("c:/tmp/check", c).Close()
	doWait()
}

func doWait() {
	done := make(chan bool)
	<-done
}

/**
*Прочитать данные из чека файла и преобразовать в объект
 */
func readCheck(fileName string) *ofd.OfdCheck {
	dat, errIo := ioutil.ReadFile(fileName)
	handleError(errIo)

	var ofdCheck ofd.OfdCheck
	err := json.Unmarshal(dat, &ofdCheck)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}

	fmt.Println("Date time:", ofdCheck.DateTime)
	for idx, item := range ofdCheck.Items {
		fmt.Println("Item", idx, ":", item)
	}
	return &ofdCheck
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
