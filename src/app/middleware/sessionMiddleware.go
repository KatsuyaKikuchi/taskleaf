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

func CheckUser(ctx *gin.Context) {
	if value, exist := ctx.Get("Session"); exist {
		if session, success := value.(*models.Session); success {
			if user, err := session.FindUser(); err != nil {
				log.Err(err)
			} else {
				ctx.Set("User", user)
			}
		}
	}
	ctx.Next()
}

func SetSessionCookie(ctx *gin.Context) {
	if value, exist := ctx.Get("Session"); exist {
		session, success := value.(*models.Session)
		if !success || session == nil {
			maxAge := -1
			// delete cookie
			ctx.SetCookie("session", "", maxAge, "/", "", true, true)
		} else {
			maxAge := 60 * 60 * 24 * 7 //!< lifetime is 1week
			ctx.SetCookie("session", session.Uuid, maxAge, "/", "", true, true)
		}
	}
	ctx.Next()
}
