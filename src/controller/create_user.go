package controller

import (
	"github.com/gin-gonic/gin"
	resterr "github.com/vhtor/metaifrn-simulados-api/src/configuration/rest_err"
)

func CreateUser(ctx *gin.Context) {
	err := resterr.NewBadRequestError("This route was called improperly")
	ctx.JSON(err.Code, err)
}