package httpserver

import (
	"net/http"
	"strconv"

	"../commonutils"
	"../persist"
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
}

func checkTotalList(context *gin.Context) {
	checkTotals, total := getCheckTotals()
	context.HTML(
		http.StatusOK,
		"checktotal.html",
		gin.H{
			"total":      total,
			"checktotal": checkTotals,
		},
	)
}

func getCheckTotals() ([]CheckTotalDto, TotalDto) {
	totals := make([]CheckTotalDto, 0)
	cashTotal := 0
	ecashTotal := 0
	for _, checkTotal := range persist.CheckTotalsList() {
		dto := CheckTotalDto{
			Id:            checkTotal.Id.String(),
			ShopId:        checkTotal.ShopId.String(),
			DateTime:      commonutils.ParseTimestamp(int64(checkTotal.DateTime)).String(),
			CashTotalSum:  strconv.Itoa(checkTotal.CashTotalSum / 100),
			EcashTotalSum: strconv.Itoa(checkTotal.EcashTotalSum / 100),
		}
		totals = append(totals, dto)
		cashTotal += checkTotal.CashTotalSum
		ecashTotal += checkTotal.EcashTotalSum
	}
	allTotal := TotalDto{
		strconv.Itoa(cashTotal / 100),
		strconv.Itoa(ecashTotal / 100),
	}
	return totals, allTotal
}
