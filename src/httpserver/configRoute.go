package httpserver

import (
	"html/template"
	"net/http"

	"../conf"
)

//Слушатель http запроса
func confList(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/config.html"))
	conf := getConfig()
	tmpl.Execute(writer, conf)
	// str, _ := json.Marshal(conf)
	// fmt.Fprintf(writer, string(str))
}

func getConfig() *conf.Configuration {
	return conf.Get()
}
