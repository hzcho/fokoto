package main

import (
	"flag"
	"fmt"
	"fokoto/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var migrationPath, configPath string

	flag.StringVar(&migrationPath, "m", "", "path to the migration dir")
	flag.StringVar(&configPath, "c", "", "path to the config dir")
	flag.Parse()

	cfg := config.MustLoad(configPath)

	if migrationPath == "" {
		panic("migration file path is empty")
	}

	fmt.Println(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName, cfg.DB.SSLMode))
	m, err := migrate.New("file://"+migrationPath,
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName, cfg.DB.SSLMode),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			panic(err)
		}
	}

	fmt.Println("migrations applied")
}
