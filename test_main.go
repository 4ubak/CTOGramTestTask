package main

import (
	"github.com/4ubak/CTOGramTestTask/internal/adapters/db/pg"
	"github.com/4ubak/CTOGramTestTask/internal/domain/core"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
	"testing"
)

var (
	app = struct {
		lg *zap.Logger
		db *pg.PostgresDb
		cr *core.St
	}{}
)

func TestMain(m *testing.M) {
	var err error

	loadConf()

	app.db, err = pg.NewPostgresDB(viper.GetString("pg_dsn"))

	if err != nil {
		log.Fatal(err)
	}

	_ = core.NewSt(app.db)
}

func loadConf() {
	confFilePath := os.Getenv("CONF_PATH")
	if confFilePath != "" {
		viper.SetConfigFile(confFilePath)

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}
	}

	// env vars are in priority
	viper.AutomaticEnv()
}

