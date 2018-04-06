package httpserver

import (
	"net/http"

	"../persistmongo"
	"github.com/gin-gonic/gin"
)

func shopList(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"shop.html",
		gin.H{
			"shops": getShops(),
		},
	)
}

func getShops() []persistmongo.Shop {
	return persistmongo.ShopList()
}
