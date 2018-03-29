package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../persist"
)

func checkTotalList(writer http.ResponseWriter, request *http.Request) {
	conf := getCheckTotals()
	str, _ := json.Marshal(conf)
	fmt.Fprintf(writer, string(str))
}

func getCheckTotals() []persist.CheckTotal {
	return persist.CheckTotalsList()
}
