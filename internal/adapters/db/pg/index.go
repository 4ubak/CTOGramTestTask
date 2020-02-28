package pg

import (
	"context"
	"github.com/4ubak/CTOGramTestTask/internal/errs"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	// ErrMsg is for ErrMsg
	ErrMsg        = "PG-error"
	dbWaitTimeout = 30 * time.Second
)

type PostgresDb struct {
	Db *sqlx.DB
}

func NewPostgresDB(dsn string) (*PostgresDb, error) {
	var err error
	if dsn == "" {
		return nil, errs.CantConnectToDb
	}

	res := &PostgresDb{}
	connectionContext, connectionContextCancel := context.WithTimeout(context.Background(), dbWaitTimeout)
	defer connectionContextCancel()

	res.Db, err = res.dbWait(connectionContext, dsn)
	if err != nil {
		return nil, err
	}

	res.Db.SetMaxOpenConns(10)
	res.Db.SetMaxIdleConns(5)
	res.Db.SetConnMaxLifetime(10 * time.Minute)

	return res, nil
}

func (pg *PostgresDb) dbWait(ctx context.Context, dsn string) (*sqlx.DB, error) {
	var err error

	var cnt uint32

	var db *sqlx.DB

	db, err = sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	for {
		err = db.PingContext(ctx)
		if err == nil || ctx.Err() != nil {
			break
		}

		time.Sleep(time.Second)
	}

	if err != nil {
		return nil, err
	}

	for {
		err = db.GetContext(ctx, &cnt, `select count(*) from calendar`)
		if err == nil || ctx.Err() != nil {
			break
		}

		time.Sleep(time.Second)
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}