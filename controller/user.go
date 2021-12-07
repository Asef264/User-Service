package controller

import (
	"net/http"
	"user/logic"
	"user/models"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	SaveNewUser(ctx *gin.Context)
	GetUserData(ctx *gin.Context)
}

type user struct {
	logic logic.UserLogicInterface
}

func NewUserHandler(logic logic.UserLogicInterface) UserHandler {

	return &user{
		logic: logic,
	}
}

func (u *user) SaveNewUser(ctx *gin.Context) {

	var NewUser models.User

	err := ctx.BindJSON(&NewUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userData, err := u.logic.SaveNewUser(ctx, NewUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, userData)

}


func(u *user)GetUserData(ctx *gin.Context) {
	
}