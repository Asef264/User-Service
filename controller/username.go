package controller

import (
	"net/http"
	"user/logic"
	"user/models"

	"github.com/gin-gonic/gin"
)

type UsernameHandller interface {
	SaveUsernameBasedUser(ctx *gin.Context)
}

type username struct {
	logic logic.UsernameLogicInterface
}

func NewLogicUsername(logic logic.UsernameLogicInterface) UsernameHandller {
	return &username{
		logic: logic,
	}
}

func (u *username) SaveUsernameBasedUser(ctx *gin.Context) {

	var usernameData models.Username

	err := ctx.BindJSON(usernameData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	username, err := u.logic.SaveUsernameBasedUser(ctx, usernameData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, username)
}
