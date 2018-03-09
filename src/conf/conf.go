package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const fileName = "conf.json"

type Configuration struct {
	IncomingCheckFolder string
}

var configuration Configuration

func Get() *Configuration {
	configuration := load()

	return configuration
}

func load() *Configuration {
	var configuration Configuration
	dat, errIo := ioutil.ReadFile(fileName)
	handleError(errIo)

	err := json.Unmarshal(dat, &configuration)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}

	return &configuration

}

/**
*Обработать ошибку
 */
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
