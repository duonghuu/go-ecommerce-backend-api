package routers

import (
	"net/http"

	"github.com/duonghuu/go-ecommerce-backend-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", Pong)
		v1.GET("/users/1", controllers.NewUserController().GetUserByID)
		v1.PUT("/ping", Pong)
		v1.POST("/ping", Pong)
		v1.DELETE("/ping", Pong)
	}

	v2 := r.Group("/api/v2")
	{
		v2.GET("/ping", Pong)
		v2.PUT("/ping", Pong)
		v2.POST("/ping", Pong)
		v2.DELETE("/ping", Pong)
	}

	return r
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong hhh",
	})
}
