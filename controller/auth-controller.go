package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raismaulana/simple_todo/dto"
	"github.com/raismaulana/simple_todo/entity"
	"github.com/raismaulana/simple_todo/helper"
	"github.com/raismaulana/simple_todo/service"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func StaticAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (controller *authController) Login(c *gin.Context) {
	var loginDTO dto.LoginDTO
	err := c.ShouldBind(&loginDTO)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to proccess request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	result := controller.authService.Login(loginDTO)
	if v, ok := result.(entity.User); ok {
		v.Token = controller.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		response := helper.BuildResponse(true, "Login Success!", v)
		c.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Login Failed!", "Username or Password is wrong!", helper.EmptyObj{})
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (controller *authController) Register(c *gin.Context) {
	var registerDTO dto.RegisterDTO
	err := c.ShouldBind(&registerDTO)

	if err != nil {
		response := helper.BuildErrorResponse("Failed to proccess request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if controller.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Register Failed!", "Email is already exists", helper.EmptyObj{})
		c.JSON(http.StatusConflict, response)
		return
	}

	user := controller.authService.Register(registerDTO)
	user.Token = controller.jwtService.GenerateToken(strconv.FormatUint(user.ID, 10))
	response := helper.BuildResponse(true, "Register Success!", user)
	c.JSON(http.StatusCreated, response)
}
