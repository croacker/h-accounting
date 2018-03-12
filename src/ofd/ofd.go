package ofd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"

	"gopkg.in/mgo.v2/bson"
)

const fileNameRegexp = "^([0-9]{2})_([0-9]{2})_([0-9]{4})_([0-9]{2})_([0-9]{2})_([0-9]{2}).*\\.json$"

//Прочитать данные из чека файла и преобразовать в объект
func ReadCheck(fileName string) (*OfdCheck, error) {
	var ofdCheck *OfdCheck
	var err error
	if isCheckFileName(fileName) {
		dat, err := ioutil.ReadFile(fileName)
		handleError(err)
		err = json.Unmarshal(dat, &ofdCheck)
		handleError(err)
	} else {
		err = errors.New("File" + fileName + "is not OFD check")
	}
	return ofdCheck, err
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
		log.Fatal(err)
	}
}

//Преобразовать в информацию о продавце(магазине)
func ToShop(check *OfdCheck) *Shop {
	return &Shop{User: check.User, UserInn: check.UserInn}
}

//Получить заголовок чека.
func ToCheckHeader(check *OfdCheck) {

}

//Тип продавец(магазин), названия полей сохранены как в оригинальном json
type Shop struct {
	Id bson.ObjectId `bson:"_id"`
	//Наименование
	User string
	//ИНН
	UserInn string
}

//Тип товар, названия полей сохранены как в оригинальном json
type Goods struct {
	Name  string
	Price int
}

//Заголовок чека
type CheckHeader struct {
}

//ОФД чек, оригинал из файла
type OfdCheck struct {
	CashTotalSum         int         `json:"cashTotalSum"`
	DateTime             int         `json:"dateTime"`
	Discount             interface{} `json:"discount"`
	DiscountSum          interface{} `json:"discountSum"`
	EcashTotalSum        int         `json:"ecashTotalSum"`
	FiscalDocumentNumber int         `json:"fiscalDocumentNumber"`
	FiscalDriveNumber    string      `json:"fiscalDriveNumber"`
	FiscalSign           int64       `json:"fiscalSign"`
	Items                []struct {
		Modifiers       interface{} `json:"modifiers"`
		Name            string      `json:"name"`
		Nds0            interface{} `json:"nds0"`
		Nds10           interface{} `json:"nds10"`
		Nds18           int         `json:"nds18"`
		NdsCalculated10 interface{} `json:"ndsCalculated10"`
		NdsCalculated18 interface{} `json:"ndsCalculated18"`
		NdsNo           interface{} `json:"ndsNo"`
		Price           int         `json:"price"`
		Quantity        float64     `json:"quantity"`
		Sum             int         `json:"sum"`
		Storno          bool        `json:"storno"`
	} `json:"items"`
	KktNumber          interface{} `json:"kktNumber"`
	KktRegID           string      `json:"kktRegId"`
	Markup             interface{} `json:"markup"`
	MarkupSum          interface{} `json:"markupSum"`
	Modifiers          interface{} `json:"modifiers"`
	Nds0               interface{} `json:"nds0"`
	Nds10              int         `json:"nds10"`
	Nds18              int         `json:"nds18"`
	NdsCalculated10    interface{} `json:"ndsCalculated10"`
	NdsCalculated18    interface{} `json:"ndsCalculated18"`
	NdsNo              interface{} `json:"ndsNo"`
	OperationType      int         `json:"operationType"`
	Operator           string      `json:"operator"`
	RequestNumber      int         `json:"requestNumber"`
	RetailPlaceAddress interface{} `json:"retailPlaceAddress"`
	ShiftNumber        int         `json:"shiftNumber"`
	StornoItems        interface{} `json:"stornoItems"`
	TaxationType       int         `json:"taxationType"`
	TotalSum           int         `json:"totalSum"`
	User               string      `json:"user"`
	UserInn            string      `json:"userInn"`
}
