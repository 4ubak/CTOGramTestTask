package interfaces

import (
	"context"
	entities "github.com/4ubak/CTOGramTestTask/internal/domain/entities"
)

type Db interface {
	Show(ctx context.Context) ([]*entities.Calendar, error)
	GetInfoByID(ctx context.Context, calendarSelect entities.CalendarSelect) (*entities.Calendar, error)
	AddEventToCalendar(ctx context.Context, calendar entities.CalendarAdd) error
	DeleteEvent(ctx context.Context, calendar entities.CalendarDelete) error
	UpdateEvent(ctx context.Context, calendar entities.Calendar) error
}