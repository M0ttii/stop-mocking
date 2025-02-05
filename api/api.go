package api

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunAPI() {
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"*"}  // Allow all headers
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = false // Must be false when AllowAllOrigins is true
	r.Use(cors.New(config))

	r.POST("/besties", BestiesHandler())
	r.GET("/followers", FollowersHandler())

	if err := r.Run(":8081"); err != nil {
		log.Fatalln("Failed to start server:", err)
	}
}
