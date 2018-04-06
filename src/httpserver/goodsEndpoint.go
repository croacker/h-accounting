package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../persist"
	"github.com/gin-gonic/gin"
)

func goods(writer http.ResponseWriter, request *http.Request) {
	goods := getGoods()
	str, _ := json.Marshal(goods)
	fmt.Fprintf(writer, string(str))
}

func goodsList(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"goods.html",
		gin.H{
			"goods": getGoods(),
		},
	)
}

func getGoods() []persist.Goods {
	return persist.GoodsList()
}
