package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// import "encoding/json"
// import "github.com/gin-gonic/gin"

func StartGin() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("./static/templates/*")
	router.GET("/", indexEndpoint)
	router.GET("/conf", confList)
	router.GET("/ofdcheck", ofdCheckList)
	router.GET("/checktotal", checkTotalList)
	router.GET("/goods", goodsList)
	router.GET("/shop", shopList)
	router.GET("/price", priceList)

	router.Run()
}

//Start запуск сервера
func StartStandart() {
	// fs := http.FileServer(http.Dir("/home/alex/projects/go/h-accounting/src/static"))
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	// http.HandleFunc("/conf", confList)
	// http.HandleFunc("/ofdcheck", ofdCheckList)
	// http.HandleFunc("/checktotal", checkTotalList)
	// http.HandleFunc("/goods", goods)
	// http.HandleFunc("/goods/list", goodsList)
	// http.HandleFunc("/price", priceList)
	// http.HandleFunc("/shop", shopList)

	http.ListenAndServe(":8081", nil)
}
