package main

import (
	_ "aisecurity/ent/dao/runtime"
	"aisecurity/middlewares"
	"aisecurity/routes"
	"aisecurity/utils"
	"aisecurity/utils/db"
	zap2 "aisecurity/utils/log/zap"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

func main() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatalf("Error loading location: %v", err)
	}
	time.Local = loc

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.InitLogger2()
	// 调用内核的Sync方法，刷新所有缓冲的日志条目。
	// 应用程序应该注意在退出之前调用Sync。
	defer func(Logger *zap2.Logger) {
		err := Logger.Sync()
		if err != nil {
			log.Fatalf("failed syncing logger: %v", err)
		}
	}(utils.Logger)

	// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	r := gin.New()
	// Set session store
	store := cookie.NewStore([]byte(os.Getenv("SESSION_KEY")))
	r.Use(sessions.Sessions("AISECURITYSESSIONID", store))
	r.Use(middlewares.RecoveryMiddleware2(), middlewares.LoggerMiddleware())
	r.Use(middlewares.Logger(), middlewares.AuditMiddleware())
	// Open the database connection
	db.InitEntClient("postgres")
	defer func() {
		err := db.EntClient.Close()
		if err != nil {
			log.Fatalf("failed closing connection to postgres: %v", err)
		}
	}()
	db.Gen()
	db.Migrate()

	// Register routes
	routes.Setup(r)

	logger, _ := zap.NewProduction()
	logger.With()

	// Run the server
	if err := r.Run("0.0.0.0:8024"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
