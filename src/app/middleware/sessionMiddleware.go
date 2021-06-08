package middleware

import (
	"app/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CheckSession(ctx *gin.Context) {
	if uuid, err := ctx.Cookie("session"); err == nil {
		session, err := models.FindSession(uuid)
		if err != nil {
			log.Warn().Msg(err.Error())
		} else {
			ctx.Set("Session", session)
		}
	} else {
		log.Warn().Msg(err.Error())
	}
	ctx.Next()
}

func SetSessionCookie(ctx *gin.Context) {
	if value, exist := ctx.Get("SessionUuid"); exist {
		if uuid, success := value.(string); success {
			maxAge := 60 * 60 * 24 * 7 //!< lifetime is 1week
			ctx.SetCookie("session", uuid, maxAge, "/", "", true, true)
		}
	}
	ctx.Next()
}
