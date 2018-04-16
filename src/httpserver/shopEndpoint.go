package httpserver

import (
	"net/http"

	"../persistsql"
	"github.com/gin-gonic/gin"
)

type ShopDto struct {
	Id      uint
	Name    string
	Address string
}

func shops(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	context.JSON(200, gin.H{
		"shops": getShops(),
	})
}

func shopsView(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"shop.html",
		gin.H{
			"shops": getShops(),
		},
	)
}

func getShops() []ShopDto {
	result := make([]ShopDto, 0)
	for _, shop := range persistsql.ShopsList() {
		dto := ShopDto{
			Id:      shop.ID,
			Name:    shop.Name,
			Address: shop.Address,
		}
		result = append(result, dto)
	}
	return result
}
