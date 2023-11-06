package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vhtor/metaifrn-simulados-api/src/model/service"
)

func NewUserControllerInterface(serviceInterface service.UserService) UserControllerInterface {
	return &userControllerInterface{serviceInterface}
}

type UserControllerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserService
}