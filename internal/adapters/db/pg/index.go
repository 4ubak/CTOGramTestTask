package pg

import (
	"database/sql"
	"fmt"
	"log"

	entities "github.com/4ubak/CTOGramTestTask/internal/domain/entities"
	_ "github.com/lib/pq"
	// c "github.com/4ubak/CTOGramTestTask/internal/constant"
)

// Show all data from db
func Show(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM calendar")
	if err != nil {
		fmt.Println("Cant find Table")
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
		return
	}

	for _, calendar := range calendars {
		fmt.Printf("ID = %d, Owner = %s, Title = %s, StartTime = %s, EndTime = %s\n", calendar.ID, calendar.Owner, calendar.Title, calendar.StartTime, calendar.EndTime)
	}
}

// //GetInfoByID ...
// func GetInfoByID(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "GET" {
// 		http.Error(w, http.StatusText(405), 405)
// 		return
// 	}

// 	id := r.FormValue("id")
// 	if id == "" {
// 		http.Error(w, http.StatusText(400), 400)
// 		return
// 	}

// 	row := db.QueryRow("SELECT * FROM calendar WHERE ID = $1", id)

// 	calendar := new(entities.Calendar)
// 	err := row.Scan(&calendar.ID, &calendar.Owner, &calendar.Title, &calendar.StartTime, &calendar.EndTime)
// 	if err == sql.ErrNoRows {
// 		http.NotFound(w, r)
// 		return
// 	} else if err != nil {
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}
// 	fmt.Printf("ID = %d, Owner = %s, Title = %s, StartTime = %s, EndTime = %s\n", calendar.ID, calendar.Owner, calendar.Title, calendar.StartTime, calendar.EndTime)
// }

// //AddEventToCalendar ...
// func AddEventToCalendar(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Error(w, http.StatusText(405), 405)
// 		return
// 	}

// 	owner := r.FormValue("Owner")
// 	title := r.FormValue("Title")
// 	startTime := r.FormValue("StartTime")
// 	endTime := r.FormValue("EndTime")

// 	fmt.Println(owner, title, startTime, endTime)

// 	if owner == "" || title == "" || startTime == "" || endTime == "" {
// 		fmt.Println("Some values not entered")
// 		http.Error(w, http.StatusText(400), 400)
// 		return
// 	}

// 	result, err := db.Exec("INSERT INTO calendar(Id, Owner, Title, StartTime, EndTime) VALUES(DEFAULT, $1, $2, $3, $4)", owner, title, startTime, endTime)

// 	if err != nil {
// 		fmt.Println("Cant insert Values")
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}

// 	rowsAffected, err := result.RowsAffected()

// 	if err != nil {
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}

// 	fmt.Fprintf(w, "Calendar %s created successfully (%d row affected)\n", owner, rowsAffected)
// }

// //DeleteEvent ...
// func DeleteEvent(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "DELETE" {
// 		fmt.Println(r.Method + ", need DELETE METHOD")
// 		http.Error(w, http.StatusText(405), 405)
// 		return
// 	}

// 	id := r.FormValue("id")
// 	result, err := db.Exec("DELETE FROM calendar WHERE Id=$1", id)

// 	if err != nil {
// 		fmt.Println("Cant Delete Values")
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}

// 	rowsAffected, err := result.RowsAffected()

// 	if err != nil {
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}

// 	fmt.Fprintf(w, "Id %s was deleted successfully (%d row affected)\n", id, rowsAffected)
// }

// //UpdateEvent ...
// func UpdateEvent(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "UPDATE" {
// 		fmt.Println(r.Method + ", need UPDATE METHOD")
// 		http.Error(w, http.StatusText(405), 405)
// 		return
// 	}

// 	id := r.FormValue("id")
// 	owner := r.FormValue("Owner")
// 	title := r.FormValue("Title")
// 	startTime := r.FormValue("StartTime")
// 	endTime := r.FormValue("EndTime")

// 	fmt.Printf("ID = %s, Owner = %s, Title = %s, StartTime = %s, EndTime = %s\n", id, owner, title, startTime, endTime)

// 	if id == "" || owner == "" || title == "" || startTime == "" || endTime == "" {
// 		fmt.Println("Some values not entered")
// 		http.Error(w, http.StatusText(400), 400)
// 		return
// 	}

// 	result, err := db.Exec("UPDATE calendar SET Owner=$2, Title=$3, StartTime=$4, EndTime=$5 WHERE Id=$1", id, owner, title, startTime, endTime)

// 	if err != nil {
// 		fmt.Println("Cant Update Values")
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}

// 	rowsAffected, err := result.RowsAffected()

// 	if err != nil {
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}

// 	fmt.Fprintf(w, "Id %s was updated successfully (%d row affected)\n", id, rowsAffected)
// }
