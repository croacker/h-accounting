package ofd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

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

/**
*Прочитать данные из чека файла и преобразовать в объект
 */
func ReadCheck(fileName string) *OfdCheck {
	dat, errIo := ioutil.ReadFile(fileName)
	handleError(errIo)

	var ofdCheck OfdCheck
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
*Обработать ошибку
 */
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
