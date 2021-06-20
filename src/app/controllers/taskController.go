package controllers

import (
	"app/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func CreateTask(ctx *gin.Context) {
	var (
		user *models.User
		err  error
	)
	if value, exist := ctx.Get("Session"); exist {
		if session, success := value.(*models.Session); success {
			if user, err = session.FindUser(); err != nil {
				log.Err(err)
			}
		}
	}

	if user != nil {
		body := ctx.PostForm("task")
		var task *models.Task
		if task, err = user.CreateTask(body); err != nil {
			log.Err(err)
		}
		log.Debug().Msgf("task = %v", task)
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Debug().Msgf(" id = %v", id)
	ctx.Redirect(http.StatusSeeOther, "/")
}
