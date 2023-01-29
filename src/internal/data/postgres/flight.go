package postgres

import (
	"ag5/internal/domain"
	"time"
)

type FlightModel struct {
	FlightId           uint64 `gorm:"primarykey"`
	FlightNo           string `gorm:"size:6"`
	ScheduledDeparture time.Time
	ScheduledArrival   time.Time
	DepartureAirport   string `gorm:"size:3"`
	ArrivalAirport     string `gorm:"size:3"`
	Status             string `gorm:"size:20"`
	AircraftCode       string `gorm:"size:3"`
	ActualDeparture    time.Time
	ActualArrival      time.Time
}

func (u *FlightModel) TableName() string {
	return "bookings.flights"
}

func (u *FlightModel) ConvertEntityToModel(flight *domain.Flight) {
	u.FlightId = flight.FlightId
	u.FlightNo = flight.FlightNo
	u.ScheduledDeparture = flight.ScheduledDeparture
	u.ScheduledArrival = flight.ScheduledArrival
	u.Status = flight.Status
}

func (u *FlightModel) ConvertModelToEntity() domain.Flight {
	return domain.Flight{
		FlightId:           u.FlightId,
		FlightNo:           u.FlightNo,
		ScheduledDeparture: u.ScheduledDeparture,
		ScheduledArrival:   u.ScheduledArrival,
		ArrivalAirport:     u.ArrivalAirport,
		DepartureAirport:   u.DepartureAirport,
		Status:             u.Status,
	}
}
