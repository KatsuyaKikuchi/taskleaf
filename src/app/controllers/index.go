package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Index(ctx *gin.Context) {
	log.Debug().Msg("Index")
}
