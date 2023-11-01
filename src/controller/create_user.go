package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	resterr "github.com/vhtor/metaifrn-simulados-api/src/configuration/rest_err"
	"github.com/vhtor/metaifrn-simulados-api/src/controller/model/request"
)

func CreateUser(ctx *gin.Context) {
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := resterr.NewBadRequestError(
			fmt.Sprintf("There are incorrect fields"),
		)
		ctx.JSON(restErr.Code, restErr)
		return
	}
	
	fmt.Println(userRequest)
}
