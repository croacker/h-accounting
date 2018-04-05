package httpserver

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"../persist"
)

func goods(writer http.ResponseWriter, request *http.Request) {
	goods := getGoods()
	str, _ := json.Marshal(goods)
	fmt.Fprintf(writer, string(str))
}

func goodsList(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/goods.html"))
	goods := getGoods()
	tmpl.Execute(writer, goods)
}

func getGoods() []persist.Goods {
	return persist.GoodsList()
}
