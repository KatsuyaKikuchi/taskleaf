package main

import (
	"github.com/gin-gonic/gin"
	zlog "github.com/rs/zerolog/log"
)

func main() {
	zlog.Info().Msg("Server Start")

	engine := gin.Default()
	engine.GET("/", routerTest)
	if err := engine.Run(":8080"); err != nil {
		zlog.Fatal().Msg(err.Error())
	}
	zlog.Info().Msg("Server End")
}
