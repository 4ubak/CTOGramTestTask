package main

import (
	"context"
	"github.com/4ubak/CTOGramTestTask/internal/adapters/db/pg"
	"github.com/4ubak/CTOGramTestTask/internal/domain/core"
	entities "github.com/4ubak/CTOGramTestTask/internal/domain/entities"
	"github.com/4ubak/CTOGramTestTask/internal/errs"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
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

	app.cr = core.NewSt(app.db)

	ec := m.Run()

	os.Exit(ec)
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

func TestShow(t *testing.T) {
	var err error

	ctx := context.Background()

	calendars, err := app.db.Show(ctx)
	require.Nil(t, err)
	require.GreaterOrEqual(t, 0, len(calendars))
}

func TestGetInfoByID(t *testing.T) {
	var err error

	ctx := context.Background()

	checkID := 1
	checkID5 := 5

	calendar, err := app.db.GetInfoByID(ctx, entities.CalendarSelect{ ID: checkID,})
	require.Equal(t, errs.IDNotFind, err)

	calendar, err := app.db.GetInfoByID(ctx, entities.CalendarSelect{ ID: checkID5,})
	require.Nil(t, err)
	//require.Equal(t, )
}

