package cmd

import (
	"fmt"
	"github.com/4ubak/CTOGramTestTask/internal/adapters/db/pg"
	"github.com/4ubak/CTOGramTestTask/internal/domain/core"
	"os"
	"os/signal"
	"syscall"

	"log"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

//Execute testing
func Execute() {
	var err error

	loadConf()
	fmt.Println(viper.GetString("pg_dsn"))
	db, err := pg.NewPostgresDB(viper.GetString("pg_dsn"))
	if err != nil {
		log.Fatal(err)
	}

	_ = core.NewSt(db)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	var exitCode int

	select {
	case <-stop:
		exitCode = 1
	}

	os.Exit(exitCode)
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