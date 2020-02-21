package cmd

import (
	"database/sql"
	"fmt"
  
	_ "github.com/lib/pq"
	"github.com/4ubak/CTOGramTestTask/internal/db/pg"
)
//Execute connect to db
func Execute() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", pg.host, pg.port, pg.user, pg.password, pg.dbname)
	db, err := sql.Open("postgres", psqlInfo)
  	if err != nil {
    panic(err)
  	}
  	defer db.Close()
  	err = db.Ping()
  	if err != nil {
    	panic(err)
  	}
  	fmt.Println("Successfully connected!")
}