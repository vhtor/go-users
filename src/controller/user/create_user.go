package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	logger "github.com/vhtor/metaifrn-simulados-api/src/configuration/log"
	"github.com/vhtor/metaifrn-simulados-api/src/configuration/validation"
	"github.com/vhtor/metaifrn-simulados-api/src/controller/model/request"
	"github.com/vhtor/metaifrn-simulados-api/src/model"
	"github.com/vhtor/metaifrn-simulados-api/src/view/convert"
	"go.uber.org/zap"
)

func (controller *userControllerInterface) CreateUser(ctx *gin.Context) {
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

	if err := controller.service.CreateUser(domain); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	logger.Info(
		"User created successfully:",
		zap.String("journey", "createUser"),
	)

	ctx.JSON(http.StatusCreated, convert.ConvertUserDomainToResponse(domain))
}
