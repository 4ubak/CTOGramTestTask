package pg

import (
	"database/sql"
	"fmt"
	entities "github.com/4ubak/CTOGramTestTask/internal/domain/entities"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// Show all data from db
func Show(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	calendars := make([]*entities.Calendar, 0)
	for rows.Next() {
		calendar := new(entities.Calendar)
		err := rows.Scan(&calendar.ID, &calendar.Owner, &calendar.Title, &calendar.StartTime, &calendar.EndTime)
		if err != nil {
			log.Fatal(err)
		}
		calendars = append(calendars, calendar)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, calendar := range calendars {
		fmt.Printf("ID = %d, Owner = %s, Title = %s, StartTime = %s, EndTime = %s\n", calendar.ID, calendar.Owner, calendar.Title, calendar.StartTime, calendar.EndTime)
	}
}
