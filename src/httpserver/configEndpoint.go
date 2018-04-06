package httpserver

import (
	"net/http"

	"../conf"
	"github.com/gin-gonic/gin"
)

//Слушатель http запроса
func confList(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"config.html",
		gin.H{
			"config": getConfig(),
		},
	)
}

func getConfig() *conf.Configuration {
	return conf.Get()
}
