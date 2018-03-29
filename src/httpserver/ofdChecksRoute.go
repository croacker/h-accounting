package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../ofd"
	"../persist"
)

func ofdCheckList(writer http.ResponseWriter, request *http.Request) {
	conf := getOfdChecks()
	str, _ := json.Marshal(conf)
	fmt.Fprintf(writer, string(str))
}

func getOfdChecks() []ofd.OfdCheck {
	return persist.OfdChecksList()
}
