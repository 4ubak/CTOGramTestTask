package pg

import (
	"database/sql"
	"errors"
	"fmt"

	// constant "github.com/4ubak/CTOGramTestTask/internal/constant"
	"log"

	entities "github.com/4ubak/CTOGramTestTask/internal/domain/entities"
	_ "github.com/lib/pq"
)

// Show all data from db
func Show(db *sql.DB) ([]*entities.Calendar, error) {
	rows, err := db.Query("SELECT * FROM calendar")
	if err != nil {
		fmt.Println("Cant find Table")
		return nil, err
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
		return nil, err
	}
	return calendars, nil
}

//GetInfoByID ...
func GetInfoByID(db *sql.DB, calendarSelect entities.CalendarSelect) (*entities.Calendar, error) {
	id := calendarSelect.ID
	row := db.QueryRow("SELECT * FROM calendar WHERE ID = $1", id)
	calendar := new(entities.Calendar)
	err := row.Scan(&calendar.ID, &calendar.Owner, &calendar.Title, &calendar.StartTime, &calendar.EndTime)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return calendar, nil
}

// AddEventToCalendar ...
func AddEventToCalendar(db *sql.DB, calendar entities.Calendar) ([]*entities.Calendar, error) {
	if calendar.Owner == "" || calendar.Title == "" || calendar.StartTime == "" || calendar.EndTime == "" {
		return nil, errors.New("Some values not entered")
	}
	result, err := db.Exec("INSERT INTO calendar(Id, Owner, Title, StartTime, EndTime) VALUES(DEFAULT, $1, $2, $3, $4)", calendar.Owner, calendar.Title, calendar.StartTime, calendar.EndTime)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	fmt.Println("Calendar %s created successfully (%d row affected)\n", calendar.Owner, rowsAffected)

	calendars, err := Show(db)
	if err != nil {
		return nil, err
	} else {
		return calendars, nil
	}
}

//DeleteEvent ...
func DeleteEvent(db *sql.DB, calendar entities.CalendarDelete) ([]*entities.Calendar, error) {
	id := calendar.ID
	result, err := db.Exec("DELETE FROM calendar WHERE Id=$1", id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	fmt.Println("Event %s successfully deleted (%d row affected)\n", calendar.ID, rowsAffected)

	calendars, err := Show(db)
	if err != nil {
		return nil, err
	} else {
		return calendars, nil
	}
}

// //UpdateEvent ...
func UpdateEvent(db *sql.DB, calendar entities.Calendar) ([]*entities.Calendar, error) {
	id := calendar.ID
	owner := calendar.Owner
	title := calendar.Title
	startTime := calendar.StartTime
	endTime := calendar.EndTime

	fmt.Printf("ID = %s, Owner = %s, Title = %s, StartTime = %s, EndTime = %s\n", id, owner, title, startTime, endTime)

	if owner == "" || title == "" || startTime == "" || endTime == "" {
		return nil, errors.New("Some values not entered")
	}

	result, err := db.Exec("UPDATE calendar SET Owner=$2, Title=$3, StartTime=$4, EndTime=$5 WHERE Id=$1", id, owner, title, startTime, endTime)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Id %s was updated successfully (%d row affected)\n", id, rowsAffected)

	calendars, err := Show(db)
	if err != nil {
		return nil, err
	} else {
		return calendars, nil
	}
}
