package httpserver

import (
	"net/http"

	"../commonutils"
	"../persistsql"
	"github.com/gin-gonic/gin"
)

type CheckTotalDto struct {
	Id            string
	ShopId        string
	DateTime      string
	CashTotalSum  string
	EcashTotalSum string
}

type TotalDto struct {
	CashTotal  string
	EcashTotal string
	AllTotal   string
}

func checkstotal(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	context.JSON(200, gin.H{
		"checkstotal": getCheckTotals(),
	})
}

func checktotalView(context *gin.Context) {
	total := getCheckTotals()
	context.HTML(
		http.StatusOK,
		"checktotal.html",
		gin.H{
			"total": total,
		},
	)
}

func getCheckTotals() TotalDto {
	cashTotal := 0
	ecashTotal := 0
	for _, checkTotal := range persistsql.ChecksList() {
		cashTotal += checkTotal.CashTotalSum
		ecashTotal += checkTotal.EcashTotalSum
	}
	totals := TotalDto{
		commonutils.ToMoneyString(cashTotal),
		commonutils.ToMoneyString(ecashTotal),
		commonutils.ToMoneyString(cashTotal + ecashTotal),
	}
	return totals
}
