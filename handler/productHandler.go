package handler

import (
	"fmt"
	"net/http"
	"tutorial/models"

	"github.com/gin-gonic/gin"
)


type Product struct {
	Id int
	Name string
}

func GetAll(context *gin.Context){
	context.AbortWithStatusJSON(http.StatusOK,[]Product{
		{
			Id: 1,
			Name: "Product 1",
		},
		{
			Id: 2,
			Name: "Product 2",
		},
	})
}

func AddProduct(context *gin.Context){
	var product Product
	if err := context.ShouldBindJSON(&product); err != nil {
		var errors []models.ErrorDetail = make([]models.ErrorDetail, 0, 1)
		errors = append(errors, models.ErrorDetail{
			ErrorType:    models.ErrorTypeValidation,
			ErrorMessage: fmt.Sprintf("%v", err),
		})
		badRequest(context, http.StatusBadRequest, "invalid request", errors)
	}
	product.Id = 10
	ok(context, http.StatusCreated,"Product Added", product)
}