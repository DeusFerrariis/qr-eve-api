package main

import (
	"fmt"
	"github.com/surrealdb/surrealdb.go"
)

func CreateEvent(event Event, conn *surrealdb.DB) (string, error) {
	// TODO: implement
	return (nil, nil)
}

func FetchEvent(ID string, conn *surrealdb.DB) (Event, error) {
	eventData, err := db.Select(fmt.Sprintf("events:%s", id))
	if err == nil {
		return (nil, err)
	}
	
	var event Event
	err := surrealdb.Unmarshal(eventData, &event) 
	if err == nil {
		return (nil, err)
	} 
	
	return (event, nil)
}

func SearchEvent(term string, start int, stop int, conn *surrealdb.DB) ([]Event, error) {
	// TODO: implement
	return (nil, nil)
}

func AttendEvent(eventID string, attendee Attendee) (string, error) {
	// TODO: implement
	return (nil, nil)
}
