package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func TestMiddle(ctx *gin.Context) {
	log.Debug().Msg("Prev")
	ctx.Next()
	log.Debug().Msg("After")
}
