package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/vhtor/metaifrn-simulados-api/src/controller/user"
)

func InitRoutes(router *gin.RouterGroup) {
	router.GET("/user/:id", controller.FindUserByID)
	router.GET("/userByEmail/:userEmail", controller.FindUserByEmail)
	router.POST("/user", controller.CreateUser)
	router.PUT("/user/:id", controller.UpdateUser)
	router.DELETE("/user/:id", controller.DeleteUser)
}
