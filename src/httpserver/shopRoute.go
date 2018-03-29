package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../persist"
)

func shopList(writer http.ResponseWriter, request *http.Request) {
	conf := getPrices()
	str, _ := json.Marshal(conf)
	fmt.Fprintf(writer, string(str))
}

func getShops() []persist.Shop {
	return persist.ShopList()
}
