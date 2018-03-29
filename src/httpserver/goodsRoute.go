package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../persist"
)

func goodsList(writer http.ResponseWriter, request *http.Request) {
	conf := getCheckTotals()
	str, _ := json.Marshal(conf)
	fmt.Fprintf(writer, string(str))
}

func getGoods() []persist.Goods {
	return persist.GoodsList()
}
