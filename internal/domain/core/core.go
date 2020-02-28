package core

import (
	"context"
	entities "github.com/4ubak/CTOGramTestTask/internal/domain/entities"
	"github.com/4ubak/CTOGramTestTask/internal/errs"
)

func (c *St) Show(ctx context.Context,) ([]*entities.Calendar, error) {
	return c.db.Show(ctx)
}

func (c *St) GetInfoByID(ctx context.Context,calendarSelect entities.CalendarSelect) (*entities.Calendar, error) {
	return c.db.GetInfoByID(ctx, calendarSelect)
}

func (c *St) AddEventToCalendar(ctx context.Context, calendar entities.CalendarAdd) error {
	owner := calendar.Owner
	title := calendar.Title
	startTime := calendar.StartTime
	endTime := calendar.EndTime
	if owner == "" || title == "" || startTime == "" || endTime == "" {
		return errs.ValuesNotFilled
	}
	return c.db.AddEventToCalendar(ctx, calendar)
}

func (c *St) DeleteEvent(ctx context.Context, calendar entities.CalendarDelete) error {
	return c.db.DeleteEvent(ctx, calendar)
}

func (c *St) UpdateEvent(ctx context.Context, calendar entities.Calendar) error {
	owner := calendar.Owner
	title := calendar.Title
	startTime := calendar.StartTime
	endTime := calendar.EndTime
	if owner == "" || title == "" || startTime == "" || endTime == "" {
		return errs.ValuesNotFilled
	}
	return c.db.UpdateEvent(ctx, calendar)
}