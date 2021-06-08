package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func Index(ctx *gin.Context) {
	user, _ := ctx.Get("LoginUser")
	log.Debug().Msgf("index user = %v", user)
	ctx.HTML(http.StatusOK, "index", gin.H{
		"title": "TaskLeaf",
		"user":  user,
	})
}
