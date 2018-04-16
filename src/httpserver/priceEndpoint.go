package httpserver

import (
	"net/http"

	"../commonutils"
	"../persistsql"
	"github.com/gin-gonic/gin"
)

type PriceDto struct {
	Id       uint
	Product  string
	ShopId   string
	DateTime string
	Price    string
}

func prices(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	context.JSON(200, gin.H{
		"prices": getPricesDto(),
	})
}

func pricesView(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"price.html",
		gin.H{
			"prices": getPricesDto(),
		},
	)
}

func getPrices() []persistsql.Price {
	return persistsql.PricesList()
}

func getPricesDto() []PriceDto {
	prices := make([]PriceDto, 0)
	for _, price := range getPrices() {
		dto := PriceDto{
			Id:       price.ID,
			Product:  price.Product.Name,
			DateTime: commonutils.ParseTimestamp(int64(price.DateTime)).String(),
			Price:    commonutils.ToMoneyString(price.Price),
		}
		prices = append(prices, dto)
	}
	return prices
}
