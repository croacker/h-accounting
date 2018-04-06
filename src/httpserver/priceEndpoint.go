package httpserver

import (
	"net/http"
	"strconv"

	"../commonutils"
	"../persist"
	"github.com/gin-gonic/gin"
)

type PriceDto struct {
	GoodsId  string
	ShopId   string
	Price    string
	DateTime string
}

func priceList(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"price.html",
		gin.H{
			"prices": getPricesDto(),
		},
	)
}

func getPrices() []persist.Price {
	return persist.PriceList()
}

func getPricesDto() []PriceDto {
	prices := make([]PriceDto, 0)
	for _, price := range getPrices() {
		dto := PriceDto{
			GoodsId:  price.Id.String(),
			ShopId:   price.ShopId.String(),
			Price:    strconv.Itoa(price.Price / 100),
			DateTime: commonutils.ParseTimestamp(int64(price.DateTime)).String(),
		}
		prices = append(prices, dto)
	}
	return prices
}
