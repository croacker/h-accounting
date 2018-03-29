package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../conf"
)

//Слушатель http запроса
func confList(writer http.ResponseWriter, request *http.Request) {
	conf := getConfig()
	str, _ := json.Marshal(conf)
	fmt.Fprintf(writer, string(str))
}

func getConfig() *conf.Configuration {
	return conf.Get()
}
