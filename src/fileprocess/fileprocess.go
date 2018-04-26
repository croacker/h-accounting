package fileprocess

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"../commonutils"
	"../ofd"
)

const backupDirName = "backup/"

//Store функция для сохранения данных
type Store func(*ofd.OfdChecks)

//ProcessFile ждать появление файла из канала
func ProcessFile(c chan string, persist Store) {
	for {
		fullPath := <-c
		time.Sleep(1 * time.Second)
		if !commonutils.IsDir(fullPath){
		fileName := filepath.Base(fullPath)
		fmt.Println("Incoming file", fileName)
		ofdChecks, err := ofd.ReadChecks(fullPath)
		if err == nil {
			persist(ofdChecks)
			toBackup(fullPath)
		} else {
			handleError(err)
		}
	}
	}
}

func toBackup(srcFile string) {
	backupDir := filepath.Dir(srcFile) + "/" + backupDirName
	commonutils.MkDirIfNotExists(backupDir)
	destinationFile := backupDir + filepath.Base(srcFile)
	os.Rename(srcFile, destinationFile)
}

//Обработать ошибку
func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
