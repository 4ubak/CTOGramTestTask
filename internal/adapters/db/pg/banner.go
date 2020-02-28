package pg

import (
	"context"
	"database/sql"
	entities "github.com/4ubak/CTOGramTestTask/internal/domain/entities"
	"github.com/4ubak/CTOGramTestTask/internal/errs"

)

// Show all data from db
func (pg *PostgresDb)Show(ctx context.Context) ([]*entities.Calendar, error) {
	rows, err := pg.Db.QueryxContext(ctx,"SELECT * FROM calendar")
	if err != nil {
		return nil, errs.TableNotExist
	}
	defer rows.Close()
	calendars := make([]*entities.Calendar, 0)
	for rows.Next() {
		calendar := new(entities.Calendar)
		err := rows.Scan(&calendar.ID, &calendar.Owner, &calendar.Title, &calendar.StartTime, &calendar.EndTime)
		if err != nil {
			return nil, errs.SqlRequestNotCorrect
		}
		calendars = append(calendars, calendar)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return calendars, nil
}

//GetInfoByID ...
func (pg *PostgresDb) GetInfoByID(ctx context.Context, calendarSelect entities.CalendarSelect) (*entities.Calendar, error) {
	id := calendarSelect.ID
	row := pg.Db.QueryRowContext(ctx,"SELECT * FROM calendar WHERE ID = $1", id)
	calendar := new(entities.Calendar)
	err := row.Scan(&calendar.ID, &calendar.Owner, &calendar.Title, &calendar.StartTime, &calendar.EndTime)
	if err == sql.ErrNoRows {
		return nil, errs.IDNotFind
	} else if err != nil {
		return nil, err
	}
	return calendar, nil
}

// AddEventToCalendar ...
func (pg *PostgresDb) AddEventToCalendar(ctx context.Context, calendar entities.CalendarAdd) (*int64, error) {
	var id *int64
	rows, err := pg.Db.QueryContext(ctx,"INSERT INTO calendar(Id, Owner, Title, StartTime, EndTime) VALUES(DEFAULT, $1, $2, $3, $4) RETURNING Id", calendar.Owner, calendar.Title, calendar.StartTime, calendar.EndTime)
	if err != nil {
		return nil, errs.IDNotFind
	}
	if rows.Next() {
		rows.Scan(&id)
	}
	return id, nil
}

//DeleteEvent ...
func (pg *PostgresDb) DeleteEvent(ctx context.Context, calendar entities.CalendarDelete) error {
	id := calendar.ID
	_, err := pg.Db.QueryxContext(ctx,"DELETE FROM calendar WHERE Id=$1", id)
	if err != nil {
		return errs.CantDeleteThisID
	}
	return nil
}

// //UpdateEvent ...
func (pg *PostgresDb) UpdateEvent(ctx context.Context, calendar entities.CalendarUpdate) error {
	id := calendar.ID
	owner := calendar.Owner
	title := calendar.Title
	startTime := calendar.StartTime
	endTime := calendar.EndTime
	_, err := pg.Db.QueryxContext(ctx,"UPDATE calendar SET Owner=$2, Title=$3, StartTime=$4, EndTime=$5 WHERE Id=$1", id, owner, title, startTime, endTime)
	if err != nil {
		return errs.CantUpdate
	}
	return nil
}
