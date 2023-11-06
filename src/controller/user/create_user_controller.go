package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	logger "github.com/vhtor/metaifrn-simulados-api/src/configuration/log"
	"github.com/vhtor/metaifrn-simulados-api/src/configuration/validation"
	"github.com/vhtor/metaifrn-simulados-api/src/controller/model/request"
	"github.com/vhtor/metaifrn-simulados-api/src/controller/model/response"
	model "github.com/vhtor/metaifrn-simulados-api/src/model"
	"github.com/vhtor/metaifrn-simulados-api/src/model/service"
	"go.uber.org/zap"
)

var (
)

func CreateUser(ctx *gin.Context) {
	logger.Info(
		"Init CreateUser controller",
		zap.String("journey", "createUser"))
	
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		logger.Error(
			"Error trying to marshal object", 
			err,
			zap.String("journey", "JSON binding"),
		)

		restErr := validation.ValidateUserError(err)

		ctx.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUser(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	service := service.NewUserService()
	if err := service.CreateUser(domain); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	response := response.UserResponse{
		ID:    "1",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info(
		"User created successfully:",
		zap.String("journey", "createUser"),
	)

	ctx.JSON(http.StatusCreated, response)
}
