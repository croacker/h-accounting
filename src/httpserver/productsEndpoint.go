package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../persistsql"
	"github.com/gin-gonic/gin"
)

type ProductDto struct {
	Id        uint
	Name      string
	Cathegory string
}

func products(writer http.ResponseWriter, request *http.Request) {
	products := getProducts()
	str, _ := json.Marshal(products)
	fmt.Fprintf(writer, string(str))
}

func productsList(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"product.html",
		gin.H{
			"products": getProducts(),
		},
	)
}

func getProducts() []ProductDto {
	result := make([]ProductDto, 0)
	for _, product := range persistsql.ProductsList() {
		dto := ProductDto{
			Id:        product.ID,
			Name:      product.Name,
			Cathegory: product.Cathegory.Name,
		}
		result = append(result, dto)
	}
	return result
}
