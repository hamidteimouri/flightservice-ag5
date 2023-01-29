package handlers

type FlightRequest struct {
	Origin      string `json:"origin" param:"origin" form:"origin"`
	Destination string `json:"destination" param:"destination" form:"destination"`
	Date        string `json:"date" param:"date" form:"date"`
}
