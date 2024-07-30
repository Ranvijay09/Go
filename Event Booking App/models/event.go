package models

import (
	"event-booking-app/db"
	"time"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	e.ID = int(id)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	dataRows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer dataRows.Close()
	var events []Event

	for dataRows.Next() {
		var event Event
		err = dataRows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(eventId int) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	dataRow := db.DB.QueryRow(query, eventId)

	var event Event
	err := dataRow.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
