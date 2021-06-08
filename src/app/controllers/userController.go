package controllers

import (
	"app/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func login(ctx *gin.Context) {

}

func Signup(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup", gin.H{
		"title": "Signup",
	})
}

func CreateAccount(ctx *gin.Context) {
	name, email, password := ctx.PostForm("name"), ctx.PostForm("email"), ctx.PostForm("password")
	user, err := models.CreateUser(name, email, password)
	if err != nil {
		// 同じEmailで登録された時とかどうする？
		log.Err(err)
		ctx.Redirect(http.StatusSeeOther, "/signup")
		return
	}
	session, err := user.CreateSession()
	if err != nil {
		log.Err(err)
	}
	log.Debug().Msgf("%v", session)
	ctx.Set("SessionUuid", session.Uuid)

	ctx.Redirect(http.StatusSeeOther, "/")
}
