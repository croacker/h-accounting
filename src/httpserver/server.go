package httpserver

import (
	"net/http"
)

// import "encoding/json"
// import "github.com/gin-gonic/gin"

//Start запуск сервера
func Start() {
	http.HandleFunc("/conf", confList)
	http.HandleFunc("/ofdcheck", ofdCheckList)
	http.HandleFunc("/checktotal", checkTotalList)
	http.HandleFunc("/goods", goodsList)
	http.HandleFunc("/price", priceList)
	http.HandleFunc("/shop", shopList)

	http.ListenAndServe(":8081", nil)
}
