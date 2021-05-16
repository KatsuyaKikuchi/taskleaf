package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func routerTest(ctx *gin.Context) {
	log.Debug().Msg("test")
}
