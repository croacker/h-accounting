package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../persist"
)

func priceList(writer http.ResponseWriter, request *http.Request) {
	conf := getPrices()
	str, _ := json.Marshal(conf)
	fmt.Fprintf(writer, string(str))
}

func getPrices() []persist.Price {
	return persist.PriceList()
}
