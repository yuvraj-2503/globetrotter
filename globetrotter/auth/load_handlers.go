package auth

import (
	"github.com/gin-gonic/gin"
	"globetrotter/auth/db"
	"globetrotter/auth/handlers"
	"globetrotter/auth/service"
	"globetrotter/common"
	"globetrotter/config"
	"log"
	token "token-manager"
)

var handler *handlers.AuthenticationHandler

func LoadHandlers(router *gin.Engine) {
	collection, err := config.Configuration.MongoConfig.GetCollection(common.AuthCollection)
	if err != nil {
		log.Fatalf("Error getting collection %s, reason: %v", common.AuthCollection, err.Error())
	}
	userDb := db.NewUserDB(collection)
	tokenManager := token.NewJwtTokenManager(config.Configuration.SecretKey)
	manager := service.NewAuthServiceImpl(userDb, tokenManager)
	handler = handlers.NewAuthenticationHandler(manager)
	loadRoutes(router)
}

func loadRoutes(router *gin.Engine) {
	group := router.Group("/api/v1/globetrotter")
	group.POST("/login", handler.Login)
	group.POST("/signup", handler.SignUp)
}
