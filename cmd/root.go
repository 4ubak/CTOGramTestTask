package cmd

import (
	"database/sql"
	"fmt"
  
	_ "github.com/lib/pq"
	"github.com/4ubak/CTOGramTestTask/internal/db/pg"
)
//Execute connect to db
func Execute() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", pg.host, pg.port, pg.ser, pg.password, pg.dbname)
}