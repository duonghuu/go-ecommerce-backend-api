package controllers

import (
	"github.com/duonghuu/go-ecommerce-backend-api/internal/service"
	"github.com/duonghuu/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	if err != nil {
		return response.ErrorResponse(c, 20003, "Email is invalid")
	}
	response.SuccessResponse(c, 20001, []string{"User data from controller"})
}
