package app

import (
	"github.com/captrep/gin-simple-crud/controller"
	"github.com/captrep/gin-simple-crud/model/domain/user"
	"github.com/captrep/gin-simple-crud/service"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	userRepository := user.NewUserRepository()
	userService := service.NewUserService(userRepository)
	UserController := controller.NewUserController(userService)
	apiRoutes(UserController)
	router.Run("127.0.0.1:3000")
}
