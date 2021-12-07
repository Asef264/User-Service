package controller

import (
	"net/http"
	"user/logic"
	"user/models"

	"github.com/gin-gonic/gin"
)

type LocationHandler interface {
	SaveLocationBasedUser(ctx *gin.Context)
}

type location struct {
	log logic.LocationBasedUserLogic
}

func NewLocationHandler(log logic.LocationBasedUserLogic) LocationHandler {
	return &location{
		log: log,
	}
}

func (l *location) SaveLocationBasedUser(ctx *gin.Context) {
	var newUserLocation models.Location

	err := ctx.BindJSON(newUserLocation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newUserByLocaion, err := l.log.SaveLocationBasedUser(ctx, newUserLocation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, newUserByLocaion)

}
