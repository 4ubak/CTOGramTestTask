package entitles

import "time"

//Calendar model
type Calendar struct {
	ID        int64
	Owner     string
	Title     string
	StartTime time.Time
	EndTime   time.Time
}
