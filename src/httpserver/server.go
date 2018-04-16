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
	router.GET("/products-view", productsView)
	router.GET("/products", products)
	router.GET("/prices-view", pricesView)
	router.GET("/prices", prices)
	router.GET("/sailers-view", sailersView)
	router.GET("/sailers", sailers)
	router.GET("/shops-view", shopsView)
	router.GET("/shops", shops)
	router.GET("/checks-view", checksView)
	router.GET("/checks", checks)
	router.GET("/checkstotal-view", checktotalView)
	router.GET("/checkstotal", checkstotal)
	router.GET("/conf-view", confList)

	router.GET("/api/products", products)

	router.Run()
}
