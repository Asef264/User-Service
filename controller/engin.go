package controller

import (
	"user/logic"
	"user/repository"

	"github.com/gin-gonic/gin"
)

func Run(port string) {
	userHandler := NewUserHandler(logic.NewLogicUserInterface(repository.NewRepoUserInterface()))
	locationHandler := NewLocationHandler(logic.NewLocBaseUserLogic(repository.NewLocBasedRepo()))
	usernameHandler := NewLogicUsername(logic.NewUsernameLogic(repository.NewRepoUsername()))
	//creating a gin engine
	engine := gin.Default()
	engine.Use(gin.Recovery())

	userPath := engine.Group("/user")
	userPath.POST("/save", userHandler.SaveNewUser)
	userPath.POST("/save", locationHandler.SaveLocationBasedUser)
	userPath.POST("/save", usernameHandler.SaveUsernameBasedUser)

	engine.Run(":" + port)
}
