package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raismaulana/simple_todo/dto"
	"github.com/raismaulana/simple_todo/helper"
	"github.com/raismaulana/simple_todo/service"
)

type UserController interface {
	UpdateUser(c *gin.Context)
	GetAllUser(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

func StaticUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (controller *userController) UpdateUser(c *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	err := c.ShouldBind(&userUpdateDTO)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to proccess request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user := controller.userService.UpdateUser(userUpdateDTO)
	response := helper.BuildResponse(true, "Updated!", user)
	c.JSON(http.StatusOK, response)
}

func (controller *userController) GetAllUser(c *gin.Context) {
	users := controller.userService.GetAll()
	response := helper.BuildResponse(true, "OK!", users)
	c.JSON(http.StatusOK, response)
}
