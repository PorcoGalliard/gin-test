package main

import (
	"log"

	"github.com/PorcoGalliard/forum/internal/configs"
	"github.com/PorcoGalliard/forum/internal/repositories/memberships"
	"github.com/PorcoGalliard/forum/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Failed initialize config")
	}

	cfg = configs.Get()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Failed initialize database")
	}

	_ = memberships.NewRepository(db)

	r.Run(cfg.Service.Port)

}