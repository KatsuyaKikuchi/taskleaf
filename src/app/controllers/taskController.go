package controllers

import (
	"app/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func CreateTask(ctx *gin.Context) {
	var user *models.User
	if value, exist := ctx.Get("User"); exist {
		user, _ = value.(*models.User)
	}

	if user != nil {
		body := ctx.PostForm("task")
		if task, err := user.CreateTask(body); err != nil {
			log.Err(err)
		} else {
			log.Debug().Msgf("task = %v", task)
		}
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}

func UpdateTask(ctx *gin.Context) {
	var (
		taskId int
		err    error
		task   *models.Task
		user   *models.User
	)
	if value, exist := ctx.Get("User"); exist {
		user, _ = value.(*models.User)
	}
	if taskId, err = strconv.Atoi(ctx.Param("id")); err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	if task, err = models.FindTask(taskId); err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	if task == nil || user == nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	if task.UserId != user.Id {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	ctx.HTML(http.StatusOK, "update_task", gin.H{
		"title": "TaskLeaf",
		"task ": task,
		"user":  user,
	})
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Debug().Msgf(" id = %v", id)
	ctx.Redirect(http.StatusSeeOther, "/")
}
