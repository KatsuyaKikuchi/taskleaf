package middleware

import (
	"github.com/gin-gonic/gin"
)

type FlashLevel int

const (
	Notice FlashLevel = iota
	Warning
	Alert
)

type Flash struct {
	Level   FlashLevel
	Message string
}

func SetFlash(ctx *gin.Context) {
	if value, exist := ctx.Get("NextFlash"); exist {
		flash, success := value.(*Flash)
		maxAge := 60 * 60 * 24
		if flash == nil || !success {
			ctx.SetCookie("Flash", "", maxAge, "/", "", true, true)
		} else {
			ctx.SetCookie("Flash", flash.Message, maxAge, "/", "", true, true)
		}
	}
}

func ClearFlash(ctx *gin.Context) {
	flash := &Flash{}
	var err error
	if flash.Message, err = ctx.Cookie("Flash"); err == nil {
		ctx.Set("Flash", flash)
	}
	maxAge := 60 * 60 * 24
	ctx.SetCookie("Flash", "", maxAge, "/", "", true, true)
}
