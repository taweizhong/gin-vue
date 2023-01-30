package main

import (
	"awesomeProject1/controller"
	"awesomeProject1/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.Use(middleware.Cors())
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/info", middleware.AuthMiddleware(), controller.Info)
}
