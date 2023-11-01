package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vhtor/metaifrn-simulados-api/src/configuration/validation"
	"github.com/vhtor/metaifrn-simulados-api/src/controller/model/request"
	"github.com/vhtor/metaifrn-simulados-api/src/controller/model/response"
)

func CreateUser(ctx *gin.Context) {
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object: %s", err.Error())
		restErr := validation.ValidateUserError(err)
		
		ctx.JSON(restErr.Code, restErr)
		return
	}
	
	fmt.Println(userRequest)
	response := response.UserResponse{
		ID: 	"1",
		Email: 	userRequest.Email,
		Name: 	userRequest.Name,
		Age: 	userRequest.Age,
	}
	ctx.JSON(http.StatusCreated, response)
}
