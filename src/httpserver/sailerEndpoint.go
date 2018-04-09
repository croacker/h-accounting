package httpserver

import (
	"net/http"

	"../persistsql"
	"github.com/gin-gonic/gin"
)

type SailerDto struct {
	Id   uint
	Name string
	Inn  string
}

func sailersView(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"sailer.html",
		gin.H{
			"sailers": getSailers(),
		},
	)
}

func getSailers() []SailerDto {
	result := make([]SailerDto, 0)
	for _, sailer := range persistsql.SailersList() {
		dto := SailerDto{
			Id:   sailer.ID,
			Name: sailer.Name,
			Inn:  sailer.Inn,
		}
		result = append(result, dto)
	}
	return result
}
