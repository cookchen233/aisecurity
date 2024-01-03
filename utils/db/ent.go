package db

import (
	"aisecurity/ent/dao"
	"aisecurity/utils"
	"context"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strings"
)

type config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Dbname   string
		Sslmode  string
	}
}

var EntClient *dao.Client

func InitEntClient(driverName string) {
	var config config
	if _, err := toml.DecodeFile("config/postgresql.toml", &config); err != nil {
		log.Printf("failed to read config/postgresql.toml, %v", err)
	}
	var err error
	EntClient, err = dao.Open(
		driverName,
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Dbname, config.Database.Password, config.Database.Sslmode),
		dao.Debug(),
		dao.Log(func(msg ...any) {
			utils.Logger.Debug(fmt.Sprint(msg))
		}))
	if err != nil {
		log.Printf("failed opening connection to postgres: %v", err)
	}
}

func Gen() {
	err := entc.Generate("./ent/schema", &gen.Config{
		Hooks:   []gen.Hook{TagFields("json")},
		Target:  "./ent/dao",
		Package: "aisecurity/ent/dao",
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
