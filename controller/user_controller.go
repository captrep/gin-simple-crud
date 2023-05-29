package controller

import (
	"log"
	"net/http"

	"github.com/captrep/gin-simple-crud/model/web"
	"github.com/captrep/gin-simple-crud/service"
	"github.com/captrep/gin-simple-crud/utils/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserControllerImpl struct {
	UserService service.UserService
}

type UserController interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	FindById(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller UserControllerImpl) CreateUser(c *gin.Context) {
	userCreateRequest := web.UserCreateRequest{}
	if err := c.ShouldBindJSON(&userCreateRequest); err != nil {
		errMsg := make(map[string][]string, 0)
		log.Println(err)
		for _, e := range err.(validator.ValidationErrors) {
			errMsg[e.StructField()] = append(errMsg[e.Field()], e.ActualTag())
		}
		c.JSON(http.StatusBadRequest, res.NewRestErr(http.StatusBadRequest, "bad request", errMsg))
		return
	}
	result, saveErr := controller.UserService.CreateUser(&userCreateRequest)
	if saveErr != nil {
		log.Println("failed create user", saveErr.Error)
		c.JSON(saveErr.Code, saveErr)
		return
	}

	c.JSON(http.StatusCreated, res.NewRestSuccess(http.StatusCreated, "created", result))
}

func (controller UserControllerImpl) GetUser(c *gin.Context) {
	users, getUsersErr := controller.UserService.GetUser()
	if getUsersErr != nil {
		c.JSON(getUsersErr.Code, getUsersErr)
		return
	}

	c.JSON(http.StatusOK, res.NewRestSuccess(http.StatusOK, "ok", users))
}

func (controller UserControllerImpl) FindById(c *gin.Context) {
	userId := c.Param("id")
	user, errFindUser := controller.UserService.FindById(&userId)

	if errFindUser != nil {
		c.JSON(errFindUser.Code, errFindUser)
		return
	}
	c.JSON(http.StatusOK, res.NewRestSuccess(http.StatusOK, "ok", user))
}

func (controller UserControllerImpl) UpdateUser(c *gin.Context) {
	userId := c.Param("id")
	userUpdateRequest := web.UserUpdateRequest{}
	if err := c.ShouldBindJSON(&userUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, res.NewRestErr(http.StatusBadRequest, "bad request", err))
		log.Println("invalid json input", err)
		return
	}

	userUpdateRequest.Id = userId

	result, errUpdateUser := controller.UserService.UpdateUser(&userUpdateRequest)
	if errUpdateUser != nil {
		c.JSON(http.StatusNotFound, errUpdateUser)
		return
	}

	c.JSON(http.StatusOK, res.NewRestSuccess(http.StatusOK, "ok", result))
}

func (controller UserControllerImpl) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	errDeleteUser := controller.UserService.DeleteUser(&userId)

	if errDeleteUser != nil {
		c.JSON(errDeleteUser.Code, errDeleteUser)
		return
	}
	c.JSON(http.StatusOK, res.NewRestSuccess(http.StatusOK, "ok", nil))
}
