package db

import (
	"aisecurity/ent"
	"context"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
	"log"
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

var EntClient *ent.Client

func InitEntClient(driverName string) {
	var config config
	if _, err := toml.DecodeFile("config/postgresql.toml", &config); err != nil {
		log.Printf("failed to read config/postgresql.toml, %v", err)
	}
	var err error
	EntClient, err = ent.Open(driverName, fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Dbname, config.Database.Password, config.Database.Sslmode))
	if err != nil {
		log.Printf("failed opening connection to postgres: %v", err)
	}
}

func Gen() {
	err := entc.Generate("./ent/schema", &gen.Config{
		Hooks: []gen.Hook{tagFields("json")},
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

// TagFields tags all fields defined in the schema with the given struct-tag.
func tagFields(name string) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				for _, field := range node.Fields {
					//field.StructTag = fmt.Sprintf("%s:%q", name, field.Name)
					field = field
				}
			}
			return next.Generate(g)
		})
	}
}

func Migration() {
	// Run the auto migration tool.
	if err := EntClient.Schema.Create(context.Background()); err != nil {
		log.Printf("failed creating schema resources: %v", err)
	}
}
