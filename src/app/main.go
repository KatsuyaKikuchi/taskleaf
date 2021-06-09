package main

import (
	"app/controllers"
	"app/middleware"
	"github.com/gin-gonic/gin"
	zlog "github.com/rs/zerolog/log"
)

func main() {
	zlog.Info().Msg("Server Start")

	engine := gin.Default()
	engine.LoadHTMLGlob("src/app/templates/*.gohtml")

	engine.GET("/", middleware.CheckSession, controllers.Index)
	engine.GET("/signup", middleware.CheckSession, controllers.Signup)
	engine.GET("/logout", middleware.CheckSession, controllers.Logout, middleware.SetSessionCookie)
	engine.GET("/login", middleware.CheckSession, controllers.Login)

	engine.POST("/create_account", controllers.CreateAccount, middleware.SetSessionCookie)
	engine.POST("/authenticate", controllers.Authenticate, middleware.SetSessionCookie)
	if err := engine.Run(":8080"); err != nil {
		zlog.Fatal().Msg(err.Error())
	}
	zlog.Info().Msg("Server End")
}
