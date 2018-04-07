package httpserver

import (
	"net/http"

	"../persistsql"
	"github.com/gin-gonic/gin"
)

type ShopDto struct {
	Id   uint
	Name string
	Inn  string
}

func shopList(context *gin.Context) {
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
			Id:   shop.ID,
			Name: shop.Name,
			Inn:  shop.Inn,
		}
		result = append(result, dto)
	}
	return result
}
