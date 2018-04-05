package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexEndpoint(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Финансы",
		},
	)
}
