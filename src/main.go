package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"time"

	"./filewatcher"
	"./ofd"
)

func printer(c chan string) {
	regex, _ := regexp.Compile("^([0-9]{2})_([0-9]{2})_([0-9]{4})_([0-9]{2})_([0-9]{2})_([0-9]{2}).*\\.json$")
	for {
		fullPath := <-c
		fileName := filepath.Base(fullPath)
		fmt.Println("File event", fileName)
		if regex.MatchString(fileName) {
			groups := regex.FindStringSubmatch(fileName)
			day := groups[1]
			month := groups[2]
			year := groups[3]
			hour := groups[4]
			minute := groups[5]
			second := groups[6]
			fmt.Println("Send day:", day, "month:",
				month, "year:", year, "hour:", hour,
				"minute", minute, "second", second)
		}
		// regex.
	}
}

func main() {
	fmt.Println(parseTimestamp(1520328120))
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
