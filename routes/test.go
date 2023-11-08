package routes

import (
  	"github.com/gin-gonic/gin"
)

import (
	"gintest/controllers"
)

func Test(r *gin.Engine) {
	r.GET("/ping", controllers.Ping)
	r.GET("/test", controllers.Test)
	r.GET("/test/:id",controllers.TestGet)
	r.PUT("/test/:id",controllers.TestPut)
	r.POST("/test",controllers.TestPost)
	r.DELETE("/test/:id",controllers.TestDelete)
}