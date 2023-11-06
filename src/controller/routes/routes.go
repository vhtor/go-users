package routes

import (
	"github.com/gin-gonic/gin"
	user_controller "github.com/vhtor/metaifrn-simulados-api/src/controller/user"
)

func InitRoutes(
	router *gin.RouterGroup, 
	userController user_controller.UserControllerInterface,
) {
	router.GET("/user/:id", userController.FindUserByID)
	router.GET("/userByEmail/:userEmail", userController.FindUserByEmail)
	router.POST("/user", userController.CreateUser)
	router.PUT("/user/:id", userController.UpdateUser)
	router.DELETE("/user/:id", userController.DeleteUser)
}
