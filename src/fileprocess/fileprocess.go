package fileprocess

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"../ofd"
)

//PersistFunc
type Store func(*ofd.OfdChecks)

//ProcessFile ждать появление файла из канала
func ProcessFile(c chan string, persist Store) {
	for {
		fullPath := <-c
		time.Sleep(1 * time.Second)
		fileName := filepath.Base(fullPath)
		fmt.Println("Incoming file", fileName)
		ofdChecks, err := ofd.ReadChecks(fullPath)
		if err == nil {
			persist(ofdChecks)
		} else {
			handleError(err)
		}
	}
}

//Обработать ошибку
func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
