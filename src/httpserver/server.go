package httpserver

import (
	// "github.com/gin-gonic/gin"
	"net/http"

	"../swagger"
)

// import "encoding/json"
// import "github.com/gin-gonic/gin"

//StartGin запуск http-сервера
func StartGin() {
	// router := gin.Default()
	// router.Static("/static", "./static")
	// router.LoadHTMLGlob("./static/templates/*")
	// router.GET("/", indexEndpoint)
	// router.GET("/products-view", productsView)
	// router.GET("/prices-view", pricesView)
	// router.GET("/sailers-view", sailersView)
	// router.GET("/shops-view", shopsView)
	// router.GET("/checks-view", checksView)
	// router.GET("/checkstotal-view", checktotalView)
	// router.GET("/conf-view", confList)

	// router.GET("/api/products", products)
	// router.GET("/api/prices", prices)
	// router.GET("/api/sailers", sailers)
	// router.GET("/api/shops", shops)
	// router.GET("/api/checks", checks)
	// router.GET("/api/checkstotal", checkstotal)

	// gin.

	// router.Run()
	// fs := http.FileServer(http.Dir("./swaggerui"))
	// http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))
	http.ListenAndServe(":8080", swagger.NewRouter())

}
