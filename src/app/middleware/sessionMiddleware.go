package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CheckSession(ctx *gin.Context) {
	log.Debug().Msg("check session")
	ctx.Next()
}
