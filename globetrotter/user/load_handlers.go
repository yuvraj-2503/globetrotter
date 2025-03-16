package user

import (
	"github.com/gin-gonic/gin"
	"globetrotter/common"
	"globetrotter/config"
	"globetrotter/user/db"
	"globetrotter/user/handlers"
	"globetrotter/user/service"
	"log"
	"user-server/auth"
)

var handler *handlers.UserHandler
var authHandler *auth.AuthHandler

func LoadHandlers(router *gin.Engine) {
	collection, err := config.Configuration.MongoConfig.GetCollection(common.UsersCollection)
	if err != nil {
		log.Fatalf("Error getting collection %s, reason: %v", common.UsersCollection, err.Error())
	}
	userDb := db.NewMongoUserStore(collection)
	manager := service.NewUserService(userDb)
	handler = handlers.NewUserHandler(manager)
	authHandler = auth.NewAuthHandler(config.Configuration.SecretKey)
	loadRoutes(router)
}

func loadRoutes(router *gin.Engine) {
	group := router.Group("/api/v1/globetrotter")
	group.Use(authHandler.Handle())
	{
		group.GET("/users", handler.RegisterUser)
		group.GET("/users/getByUsername", handler.GetUserByUsername)
		group.GET("/users/myScore", handler.GetUserById)
	}
}
