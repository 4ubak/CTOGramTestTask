package entitles

//Calendar ...
type Calendar struct {
	ID        int64
	Owner     string
	Title     string
	StartTime string
	EndTime   string
}

//CalendarSelect ...
type CalendarSelect struct {
	ID int64
}

//CalendarDelete ...
type CalendarDelete struct {
	ID int64
}

//CalendarUpdate ...
type CalendarUpdate struct {
	ID        int64
	Owner     string
	Title     string
	StartTime string
	EndTime   string
}
