package ofd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
)

const fileNameRegexp = "^([0-9]{2})_([0-9]{2})_([0-9]{4})_([0-9]{2})_([0-9]{2})_([0-9]{2}).*\\.json$"

//Прочитать данные из чека файла и преобразовать в объект
func ReadChecks(fileName string) (*OfdChecks, error) {
	var ofdChecks *OfdChecks
	var err error
	if isCheckFileName(fileName) {
		dat, err := ioutil.ReadFile(fileName)
		handleError(err)
		err = json.Unmarshal(dat, &ofdChecks)
		handleError(err)
		if err != nil {
			var ofdCheck *OfdCheck
			err = json.Unmarshal(dat, &ofdCheck)
			ofdChecks = &OfdChecks{
				*ofdCheck,
			}
		}
	} else {
		err = errors.New("File" + fileName + "is not OFD check")
	}
	return ofdChecks, err
}

//Проверить наименование файла.
func isCheckFileName(fullPath string) bool {
	regex, _ := regexp.Compile(fileNameRegexp)
	fileName := filepath.Base(fullPath)
	return regex.MatchString(fileName)
}

//Разделить имя файла на составляющие.
func printFileNameDetails(fullPath string) {
	regex, _ := regexp.Compile(fileNameRegexp)
	fileName := filepath.Base(fullPath)
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
}

//Обработать ошибку.
func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}

//Получить заголовок чека.
func ToCheckHeader(check *OfdCheck) {

}

//Заголовок чека
type CheckHeader struct {
}
