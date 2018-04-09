package commonutils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

//Текущий каталог
func CurrentFolder() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current folder:", dir)
	return dir
}

//Преобразовать timestamp из чека(например ofdCheck.DateTime) во время.
func ParseTimestamp(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

func ToMoneyString(v int) string {
	floatMoney := float64(v) / 100
	return fmt.Sprintf(fmt.Sprintf("%%.%df", 2), floatMoney)
}

func ToDatetimeString(timestamp int) string {
	return ParseTimestamp(int64(timestamp)).Format("02-01-2006 15:04:05")
}
