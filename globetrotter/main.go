package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"globetrotter/config"
	"globetrotter/game"
	"globetrotter/invite"
	"globetrotter/user"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()
	ctx := context.TODO()

	user.LoadDB(&ctx)
	user.LoadHandlers(router)

	game.LoadDB(&ctx)
	game.LoadHandlers(router)

	invite.LoadHandlers(router)

	public := router.Group("/api/v1")
	public.GET("/health", health)

	err := router.Run("0.0.0.0:" + strconv.Itoa(config.Configuration.ServerPort))
	if err != nil {
		log.Panicf("Failed to start globetrotter server, reason: %v", err)
		return
	}

	log.Println("Started globetrotter server at port: " + strconv.Itoa(config.Configuration.ServerPort))
}

func health(c *gin.Context) {
	c.Data(http.StatusOK, gin.MIMEPlain, []byte{'0'})
}
