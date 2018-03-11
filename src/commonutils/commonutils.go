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
