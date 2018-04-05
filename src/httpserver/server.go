package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// import "encoding/json"
// import "github.com/gin-gonic/gin"

func StartGin() {
	router := gin.Default()
	router.LoadHTMLGlob("./static/templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})
	router.Run()
}

//Start запуск сервера
func StartStandart() {
	// fs := http.FileServer(http.Dir("/home/alex/projects/go/h-accounting/src/static"))
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/conf", confList)
	http.HandleFunc("/ofdcheck", ofdCheckList)
	http.HandleFunc("/checktotal", checkTotalList)
	http.HandleFunc("/goods", goods)
	http.HandleFunc("/goods/list", goodsList)
	http.HandleFunc("/price", priceList)
	http.HandleFunc("/shop", shopList)

	http.ListenAndServe(":8081", nil)
}
