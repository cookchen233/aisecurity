package main

import (
	"aisecurity/middlewares"
	"aisecurity/routes"
	"aisecurity/utils/db"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Printf("Error loading location: %v", err)
	}
	time.Local = loc

	// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()

	r.Use(middlewares.Logger())

	// Register routes
	routes.Setup(r)

	// Open the database connection
	db.InitEntClient("postgres")
	defer func() {
		err := db.EntClient.Close()
		if err != nil {
			log.Printf("failed closing connection to postgres: %v", err)
		}
	}()
	db.Gen()
	db.Migration()

	// Run the server
	if err := r.Run("0.0.0.0:8024"); err != nil {
		log.Printf("Failed to start server: %v", err)
	}
}
