package invite

import (
	"github.com/gin-gonic/gin"
	"globetrotter/common"
	"globetrotter/config"
	"globetrotter/invite/db"
	"globetrotter/invite/handlers"
	"globetrotter/invite/service"
	userDb "globetrotter/user/db"
	user "globetrotter/user/service"
	"log"
	"user-server/auth"
)

var inviteHandler *handlers.InviteHandler
var authHandler *auth.AuthHandler

func LoadHandlers(router *gin.Engine) {
	collection, err := config.Configuration.MongoConfig.GetCollection(common.InvitesCollection)
	if err != nil {
		log.Fatalf("Error getting collection %s, reason: %v", common.InvitesCollection, err.Error())
	}

	userColl, err := config.Configuration.MongoConfig.GetCollection(common.UsersCollection)
	if err != nil {
		log.Fatalf("Error getting collection %s, reason: %v", common.UsersCollection, err.Error())
	}

	inviteDB := db.NewInviteDB(collection)
	userDB := userDb.NewMongoUserStore(userColl)
	userService := user.NewUserService(userDB)
	inviteService := service.NewInviteService(inviteDB, userService)
	inviteHandler = handlers.NewInviteHandler(inviteService)
	authHandler = auth.NewAuthHandler(config.Configuration.SecretKey)
	loadRoutes(router)
}

func loadRoutes(router *gin.Engine) {
	group := router.Group("/api/v1/globetrotter")
	group.Use(authHandler.Handle())
	{
		group.GET("/invite", inviteHandler.GetInviteLink)
		group.GET("/inviter/score", inviteHandler.GetInviterScore)
	}
}
