package controllers

import (
	"app/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func Index(ctx *gin.Context) {
	var (
		user *models.User
		err  error
	)
	if value, exist := ctx.Get("Session"); exist {
		if session, ok := value.(*models.Session); ok {
			user, err = session.FindUser()
			if err != nil {
				log.Err(err)
			}
		}
	}
	tasks := make([]models.Task, 0)
	if user != nil {
		if tasks, err = user.Tasks(); err != nil {
			log.Err(err)
		}
	}

	ctx.HTML(http.StatusOK, "index", gin.H{
		"title": "TaskLeaf",
		"user":  user,
		"tasks": tasks,
	})
}
