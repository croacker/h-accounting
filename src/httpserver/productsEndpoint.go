package httpserver

import (
	"net/http"

	"../persistsql"
	"github.com/gin-gonic/gin"
)

type ProductDto struct {
	Id        uint
	Name      string
	Cathegory string
}

func products(context *gin.Context) {
	products := getProducts()
	// str, _ := json.Marshal(products)
	context.JSON(200, gin.H{
		"products": products,
	})
	// fmt.Fprintf(writer, string(str))
}

func productsView(context *gin.Context) {
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
