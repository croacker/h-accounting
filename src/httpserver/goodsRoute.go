package httpserver

import (
	"html/template"
	"net/http"

	"../persist"
)

func goodsList(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/goods.html"))
	goods := getGoods()
	tmpl.Execute(writer, goods)
	// str, _ := json.Marshal(goods)
	// fmt.Fprintf(writer, string(str))
}

func getGoods() []persist.Goods {
	return persist.GoodsList()
}
