package test

import (
	"testing"

	"github.com/4ubak/CTOGramTestTask/cmd"
	"github.com/4ubak/CTOGramTestTask/internal/adapters/db/pg"
	entitles "github.com/4ubak/CTOGramTestTask/internal/domain/entities"
)

func Test_Start(t *testing.T) {
	Test_CheckConnect(t)
	Test_Show(t)
	Test_Select(t)
	Test_Add(t)
	Test_Delete(t)
	Test_Update(t)
}

func Test_CheckConnect(t *testing.T) {
	value := cmd.Execute()
	if value == "Successfully connected!" {
		t.Logf("Success")
	} else {
		t.Errorf("Failed %s", value)
	}
}

func Test_Show(t *testing.T) {
	calendars, err := pg.Show(cmd.Db)
	if err != nil {
		t.Errorf("Failed %s", err)
	} else {
		for _, calendar := range calendars {
			t.Logf("ID = %d, Owner = %s, Title = %s, StartTime = %s, EndTime = %s\n", calendar.ID, calendar.Owner, calendar.Title, calendar.StartTime, calendar.EndTime)
		}
	}
}

func Test_Select(t *testing.T) {
	calendar := entitles.CalendarSelect{
		ID: 3,
	}
	id, err := pg.GetInfoByID(cmd.Db, calendar)
	if err != nil {
		t.Errorf("Failed %s", err)
	} else {
		t.Logf("%v", id)
	}
}

func Test_Add(t *testing.T) {
	calendar := entitles.Calendar{
		Owner:     "Test",
		Title:     "TestTitle",
		StartTime: "27.02.2020",
		EndTime:   "28.02.2020",
	}
	id, err := pg.AddEventToCalendar(cmd.Db, calendar)
	if err != nil {
		t.Errorf("Failed %s", err)
	} else {
		t.Logf("%v", id)
	}
}

func Test_Delete(t *testing.T) {
	calendar := entitles.CalendarDelete{
		ID: 1,
	}
	id, err := pg.DeleteEvent(cmd.Db, calendar)
	if err != nil {
		t.Errorf("Failed %s", err)
	} else {
		t.Logf("%v", id)
	}
}

func Test_Update(t *testing.T) {
	calendar := entitles.Calendar{
		ID:        1,
		Owner:     "Test",
		Title:     "Update",
		StartTime: "27.02.2020",
		EndTime:   "28.02.2020",
	}
	id, err := pg.UpdateEvent(cmd.Db, calendar)
	if err != nil {
		t.Errorf("Failed %s", err)
	} else {
		t.Logf("%v", id)
	}
}
