package main

import (
	"context"
	"github.com/4ubak/CTOGramTestTask/internal/adapters/db/pg"
	"github.com/4ubak/CTOGramTestTask/internal/domain/core"
	entities "github.com/4ubak/CTOGramTestTask/internal/domain/entities"
	"github.com/4ubak/CTOGramTestTask/internal/errs"
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

	app.db, err = pg.NewPostgresDB("postgres://postgres:your-password@localhost:5432/calendar_demo")

	if err != nil {
		log.Fatal(err)
	}

	app.cr = core.NewSt(app.db)

	ec := m.Run()

	os.Exit(ec)
}

func TestShow(t *testing.T) {
	var err error

	ctx := context.Background()

	calendars, err := app.cr.Show(ctx)
	require.Nil(t, err)
	require.LessOrEqual(t, 0, len(calendars))
}

func TestGetInfoByID(t *testing.T) {
	var err error

	ctx := context.Background()

	var checkID int64 = 1
	var checkID5 int64 = 5

	_, err = app.cr.GetInfoByID(ctx, entities.CalendarSelect{ ID: checkID,})
	require.Equal(t, errs.IDNotFind, err)

	var testCalendar = entities.Calendar{
		ID: checkID5,
		Owner: "Test",
		Title: "TestTitle",
		StartTime: "27.02.2020",
		EndTime: "28.02.2020",
	}
	calendar, err := app.db.GetInfoByID(ctx, entities.CalendarSelect{ ID: checkID5,})
	require.Nil(t, err)
	require.Equal(t, testCalendar.ID, calendar.ID)
	require.Equal(t, testCalendar.Owner, calendar.Owner)
	require.Equal(t, testCalendar.Title, calendar.Title)
	require.Equal(t, testCalendar.StartTime, calendar.StartTime)
	require.Equal(t, testCalendar.EndTime, calendar.EndTime)
}

func TestAddValue(t *testing.T) {
	var err error

	var idUser int64 = 29 //increase after each test +2
	ctx := context.Background()

	id, err := app.cr.AddEventToCalendar(ctx, entities.CalendarAdd{
		Owner:     "Tim",
		Title:     "TestAddValue",
		StartTime: "28.02.2020",
		EndTime:   "28.02.2020",
	})
	require.Nil(t, err)
	require.Equal(t, &idUser, id)

	id, err = app.cr.AddEventToCalendar(ctx, entities.CalendarAdd{
		Owner:     "",
		Title:     "TestAddValue",
		StartTime: "28.02.2020",
		EndTime:   "28.02.2020",
	})
	require.Nil(t, id)
	require.Equal(t, errs.ValuesNotFilled, err)
}

func TestDeleteValue(t *testing.T) {
	var err error

	ctx := context.Background()

	var idUser int64 = 29 //decrease after each test -1
	err = app.cr.DeleteEvent(ctx, entities.CalendarDelete{ID: idUser,})
	require.Nil(t, err)

	var idUser2 int64 = 0
	err = app.cr.DeleteEvent(ctx, entities.CalendarDelete{ID: idUser2,})
	require.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	var err error

	ctx := context.Background()

	var idUser int64 = 7 //increase after each test +1
	calendarUpdate := entities.CalendarUpdate{
		ID:        idUser,
		Owner:     "Andrey",
		Title:     "Updated",
		StartTime: "29.02.2020",
		EndTime:   "29.02.2020",
	}
	err = app.cr.UpdateEvent(ctx, calendarUpdate)
	require.Nil(t, err)

	calendarNotFilled := entities.CalendarUpdate{
		ID:        0,
		Owner:     "",
		Title:     "Updated",
		StartTime: "29.02.2020",
		EndTime:   "29.02.2020",
	}
	err = app.cr.UpdateEvent(ctx, calendarNotFilled)
	require.Equal(t, errs.ValuesNotFilled, err)
}
