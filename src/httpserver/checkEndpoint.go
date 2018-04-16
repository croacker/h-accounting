package httpserver

import (
	"net/http"

	"../commonutils"
	"../persistsql"
	"github.com/gin-gonic/gin"
)

type CheckDto struct {
	Id            uint
	DateTime      string
	Shop          string
	CashTotalSum  string
	EcashTotalSum string
}

func checks(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	context.JSON(200, gin.H{
		"checks": getChecksDto(),
	})
}

func checksView(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"check.html",
		gin.H{
			"checks": getChecksDto(),
		},
	)
}

func getChecks() []persistsql.CheckHeader {
	return persistsql.ChecksList()
}

func getChecksDto() []CheckDto {
	prices := make([]CheckDto, 0)
	for _, check := range getChecks() {
		dto := CheckDto{
			Id:            check.ID,
			DateTime:      commonutils.ToDatetimeString(check.DateTime),
			Shop:          check.Shop.Name,
			CashTotalSum:  commonutils.ToMoneyString(check.CashTotalSum),
			EcashTotalSum: commonutils.ToMoneyString(check.EcashTotalSum),
		}
		prices = append(prices, dto)
	}
	return prices
}
