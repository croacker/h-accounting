package httpserver

import (
	"html/template"
	"net/http"

	"../persist"
)

func shopList(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/shop.html"))
	shops := getShops()
	tmpl.Execute(writer, shops)
	// str, _ := json.Marshal(conf)
	// fmt.Fprintf(writer, string(str))
}

func getShops() []persist.Shop {
	return persist.ShopList()
}
