package httpserver

import (
	"github.com/gin-gonic/gin"
)

// import "encoding/json"
// import "github.com/gin-gonic/gin"

//StartGin запуск http-сервера
func StartGin() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("./static/templates/*")
	router.GET("/", indexEndpoint)
	router.GET("/productsView", productsView)
	router.GET("/products", products)
	router.GET("/pricesView", pricesView)
	router.GET("/sailersView", sailersView)
	router.GET("/shopsView", shopsView)
	router.GET("/checksView", checksView)
	router.GET("/checktotalView", checktotalView)
	router.GET("/confView", confList)

	router.Run()
}
