package httpserver

import (
	"net/http"

	"../ofd"
	"../persistmongo"
	"github.com/gin-gonic/gin"
)

func ofdCheckList(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"ofdcheck.html",
		gin.H{
			"ofdchecks": getOfdChecks(),
		},
	)
}

func getOfdChecks() []ofd.OfdCheck {
	return persistmongo.OfdChecksList()
}
