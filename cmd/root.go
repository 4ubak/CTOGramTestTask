package cmd

import (
	"database/sql"
	"fmt"
  
	_ "github.com/lib/pq"
	"github.com/4ubak/CTOGramTestTask/internal/adapters/db/pg"
	"github.com/4ubak/CTOGramTestTask/internal/domain/entities"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "calendar_demo"
)

//Execute connect to db
func Execute() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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
  
	sqlStatement := `SELECT * FROM users WHERE ID=$1;`
	var calendar entitles.Calendar
	row := db.QueryRow(sqlStatement, 1)
	err := row.Scan(&calendar.ID, &calendar.Owner, &calendar.Title, &calendar.StartTime, &calendar.EndTime)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(user)
	default:
		panic(err)
	}
}