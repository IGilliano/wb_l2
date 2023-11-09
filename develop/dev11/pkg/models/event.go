package models

import "time"

type Event struct {
	EventId     int       `json:"event_id"`
	UserId      int       `json:"user_id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

//
//type EventCalendar struct {
//	Event  Event `json:"event"`
//}
