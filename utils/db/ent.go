package db

import (
	"aisecurity/ent/dao"
	"aisecurity/utils"
	"context"
	sql2 "database/sql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"log"
	"os"
	"strings"
	"time"
)

type config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string `toml:"db_name"`
		SSLMode  string `toml:"ssl_mode"`
	}
}

var EntClient *dao.Client
var DB *sql2.DB

func InitEntClient(driverName string) {
	var dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSL_MODE"),
	)
	var err error
	EntClient, err = dao.Open(
		driverName,
		dsn,
		dao.Debug(),
		dao.Log(func(msg ...any) {
			if gin.Mode() == "debug" {
				utils.Logger.Debug(fmt.Sprint(msg))
			}
		}))
	if err != nil {
		log.Printf("failed opening connection to postgres: %v", err)
	}
	// 相比于直接Open DSN获取连接，封装entgo.io/ent/dialect/sql.*DB显然更加适合生产环境，即有连接池的支持，又可以支持执行裸SQL
	//drv, err := sql.Open(driverName, dsn)
	//if err != nil {
	//	utils.Logger.Error(err.Error())
	//	os.Exit(1)
	//}
	//
	//db := drv.DB()
	//db.SetMaxIdleConns(10)
	//db.SetMaxOpenConns(100)
	//db.SetConnMaxLifetime(time.Hour)
	//EntClient = dao.NewClient(dao.Driver(drv))
}

func Generate() {
	err := entc.Generate("./ent/schema", &gen.Config{
		Hooks:   []gen.Hook{TagFields("json")},
		Target:  "./ent/dao",
		Package: "aisecurity/ent/dao",
		Features: []gen.Feature{
			gen.FeatureEntQL,
			gen.FeatureExecQuery,
			gen.FeaturePrivacy,
		},
	}, []entc.Option{
		//entc.TemplateDir("./ent/templates"),
	}...)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

// TagFields tags all fields defined in the schema with the given struct-tag.
func TagFields(name string) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				for _, field := range node.Fields {
					field.StructTag = strings.Replace(field.StructTag, ",omitempty", "", 1)
				}
			}
			return next.Generate(g)
		})
	}
}

func Migrate() {
	ctx := context.Background()
	filename := fmt.Sprintf("migration/%s.sql", time.Now().Format("2006-01-02"))

	// Open the file for appending, create it if it does not exist
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		utils.Logger.Error("failed to create or open migrate file: %v", zap.Error(err))
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			utils.Logger.Error("failed closing migrate file", zap.Error(err))
		}
	}()

	// Write schema to the file
	if err := EntClient.Schema.WriteTo(ctx, f); err != nil {
		utils.Logger.Error("failed printing schema changes", zap.Error(err))
	} else {
		if err := EntClient.Schema.Create(ctx); err != nil {
			utils.Logger.Error("failed creating schema resources", zap.Error(err))
		}
	}
}
