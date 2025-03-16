package game

import (
	"github.com/gin-gonic/gin"
	"globetrotter/common"
	"globetrotter/config"
	"globetrotter/game/db"
	"globetrotter/game/handlers"
	"globetrotter/game/service"
	userDb "globetrotter/user/db"
	userService "globetrotter/user/service"
	"log"
	"user-server/auth"
)

var handler *handlers.GameHandler
var authHandler *auth.AuthHandler

func LoadHandlers(router *gin.Engine) {
	collection, err := config.Configuration.MongoConfig.GetCollection(common.DestinationCollection)
	if err != nil {
		log.Fatalf("Error getting collection %s, reason: %v", common.DestinationCollection, err.Error())
	}
	userColl, err := config.Configuration.MongoConfig.GetCollection(common.UsersCollection)
	if err != nil {
		log.Fatalf("Error getting collection %s, reason: %v", common.UsersCollection, err.Error())
	}
	destDB := db.NewDestinationDBStore(collection)
	userDB := userDb.NewMongoUserStore(userColl)
	userservice := userService.NewServerService(userDB)
	manager := service.NewGameServiceImpl(destDB, userservice)
	handler = handlers.NewGameHandler(manager)
	authHandler = auth.NewAuthHandler(config.Configuration.SecretKey)
	loadRoutes(router)
}

func loadRoutes(router *gin.Engine) {
	group := router.Group("/api/v1/globetrotter")
	group.Use(authHandler.Handle())
	{
		group.GET("/question", handler.GetRandomQuestion)
		group.POST("/answer", handler.SubmitAnswer)
	}
}
