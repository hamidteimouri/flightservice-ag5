package domain

import (
	"time"
)

type Flight struct {
	FlightId           uint64    `json:"flight_id"`
	FlightNo           string    `json:"flight_no"`
	ScheduledDeparture time.Time `json:"scheduled_departure"`
	ScheduledArrival   time.Time `json:"scheduled_arrival"`
	ArrivalAirport     string    `json:"arrival_airport"`
	DepartureAirport   string    `json:"departure_airport"`
	Status             string    `json:"status"`
}

type FlightSearchRequest struct {
	Origin      string
	Destination string
	Date        time.Time
}
