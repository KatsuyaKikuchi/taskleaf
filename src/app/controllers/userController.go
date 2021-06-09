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

func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login", gin.H{
		"title": "Login",
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
	ctx.Set("Session", session)

	ctx.Redirect(http.StatusSeeOther, "/")
}

func Authenticate(ctx *gin.Context) {
	email, password := ctx.PostForm("email"), ctx.PostForm("password")
	user, err := models.FindUser(email, password)
	if err != nil {
		log.Err(err)
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}
	session, err := user.CreateSession()
	if err != nil {
		log.Err(err)
	}
	ctx.Set("Session", session)
	ctx.Redirect(http.StatusSeeOther, "/")
}

func Logout(ctx *gin.Context) {
	if value, exist := ctx.Get("Session"); exist {
		if session, ok := value.(*models.Session); ok {
			if err := session.DeleteSession(); err != nil {
				log.Err(err)
			}
		}
	}
	ctx.Set("Session", nil)
	ctx.Redirect(http.StatusSeeOther, "/")
}
