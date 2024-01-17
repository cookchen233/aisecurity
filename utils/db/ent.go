package db

import (
	"aisecurity/ent/dao"
	"aisecurity/utils"
	"context"
	sql2 "database/sql"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
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
	var config config
	if _, err := toml.DecodeFile("config/postgresql.toml", &config); err != nil {
		log.Printf("failed to read config/postgresql.toml, %v", err)
	}
	var dsn = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", config.Database.Host, config.Database.Port, config.Database.User, config.Database.DBName, config.Database.Password, config.Database.SSLMode)
	//var err error
	//EntClient, err = dao.Open(
	//	driverName,
	//    entDsn,
	//	dao.Debug(),
	//	dao.Log(func(msg ...any) {
	//		if gin.Mode() == "debug" {
	//			utils.Logger.Debug(fmt.Sprint(msg))
	//		}
	//	}))
	//if err != nil {
	//	log.Printf("failed opening connection to postgres: %v", err)
	//}
	//var dbDsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/<%s>?parseTime=True&sslmode=%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DBName, config.Database.SSLMode)
	drv, err := sql.Open(driverName, dsn)
	if err != nil {
		utils.Logger.Error(err.Error())
		os.Exit(1)
	}

	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// 相比于直接Open DSN获取连接，封装entgo.io/ent/dialect/sql.*DB显然更加适合生产环境，即有连接池的支持，又可以支持执行裸SQL
	EntClient = dao.NewClient(dao.Driver(drv))
}

func Gen() {
	err := entc.Generate("./ent/schema", &gen.Config{
		Hooks:   []gen.Hook{TagFields("json")},
		Target:  "./ent/dao",
		Package: "aisecurity/ent/dao",
		Features: []gen.Feature{
			gen.FeatureEntQL,
			gen.FeatureExecQuery,
			gen.FeaturePrivacy,
		},
	})
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
	// Dump migration changes to an SQL script.
	f, err := os.Create("migrate.sql")
	if err != nil {
		log.Printf("create migrate file: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Printf("failed closing migrate file: %v", err)
		}
	}(f)
	if err := EntClient.Schema.WriteTo(ctx, f); err != nil {
		log.Printf("failed printing schema changes: %v", err)
	} else {
		if err := EntClient.Schema.Create(ctx); err != nil {
			log.Printf("failed creating schema resources: %v", err)
		}
	}
}
