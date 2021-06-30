package handler

import (
	"fmt"
	"net/http"
	"time"
	"tutorial/models"
	"tutorial/token"

	"github.com/gin-gonic/gin"
)

func LoginHandler(context *gin.Context) {
	var loginObj models.LoginRequest
	if err := context.ShouldBindJSON(&loginObj); err != nil {
		var errors []models.ErrorDetail = make([]models.ErrorDetail, 0, 1)
		errors = append(errors, models.ErrorDetail{
			ErrorType:    models.ErrorTypeValidation,
			ErrorMessage: fmt.Sprintf("%v", err),
		})
		badRequest(context, http.StatusBadRequest, "invalid request", errors)
	}
	// validate the loginObj for valid credential adn if these are valid then

	var claims = &models.JwtClaims{}
	claims.ComapnyId = "ComapnyId"
	claims.Username = loginObj.UserName
	claims.Roles = []int{1,2,3}
	claims.Audience = context.Request.Header.Get("Referer") // get it from Referer header

	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(2) * time.Hour)
	tokeString, err := token.GenrateToken(claims, expirationTime)

	if err != nil {
		badRequest(context, http.StatusBadRequest, "error in gerating token", []models.ErrorDetail{
			{
				ErrorType:    models.ErrorTypeError,
				ErrorMessage: err.Error(),
			},
		})
	}

	ok(context, http.StatusOK, "token created", tokeString)
}
