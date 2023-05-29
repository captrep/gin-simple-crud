package app

import "github.com/captrep/gin-simple-crud/controller"

func apiRoutes(userController controller.UserController) {
	router.POST("users", userController.CreateUser)
	router.GET("users", userController.GetUser)
	router.GET("users/:id", userController.FindById)
	router.PUT("users/:id", userController.UpdateUser)
	router.DELETE("users/:id", userController.DeleteUser)
}
