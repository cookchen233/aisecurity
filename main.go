package main

import (
	_ "aisecurity/ent/dao/runtime"
	"aisecurity/middlewares"
	"aisecurity/routes"
	"aisecurity/services"
	"aisecurity/utils"
	"aisecurity/utils/db"
	zap2 "aisecurity/utils/log/zap"
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

func InitMain() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatalf("Error loading location: %v", err)
	}
	time.Local = loc

	// load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init logging
	utils.InitLogger2()

	// Open the database connection
	db.InitEntClient("postgres")

}
func main() {
	InitMain()

	defer func(Logger *zap2.Logger) {
		err := Logger.Sync()
		if err != nil {
			log.Fatalf("failed syncing logger: %v", err)
		}
	}(utils.Logger)

	defer func() {
		err := db.EntClient.Close()
		if err != nil {
			log.Fatalf("failed closing connection to postgres: %v", err)
		}
	}()

	logger, _ := zap.NewProduction()
	logger.With()

	utils.InitWechatOfficialAccount()
	services.NewSweepScheduleService(context.Background()).SetCrontab()

	// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	r := gin.New()

	// Set session store
	store := cookie.NewStore([]byte(os.Getenv("SESSION_KEY")))
	r.Use(sessions.Sessions("AISECURITYSESSIONID", store))

	// Register middlewares
	r.Use(
		middlewares.Recovery(),
		middlewares.RequestLog(),
		//middlewares.JoyRequestLog(),
		middlewares.CORSMiddleware(),
	)

	// Register routes
	routes.Setup(r)

	// Run the server
	if err := r.Run("0.0.0.0:8024"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
