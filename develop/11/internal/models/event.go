package models

import "time"

type Event struct {
	ID   int       `json:"event_id"`
	Date time.Time `json:"event_date"`
	Name string    `json:"event_name"`
}
