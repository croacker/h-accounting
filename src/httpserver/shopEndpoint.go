package httpserver

import (
	"net/http"

	"../persist"
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

func getShops() []persist.Shop {
	return persist.ShopList()
}
