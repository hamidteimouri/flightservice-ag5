package handlers

import (
	"ag5/internal/domain"
	"github.com/labstack/echo/v4"
	"time"
)

type FlightHandler struct {
	ctrl *domain.FlightController
}

func NewFlightHandler(ctrl *domain.FlightController) *FlightHandler {
	return &FlightHandler{ctrl: ctrl}
}

func (f *FlightHandler) BestMatch(c echo.Context) error {

	req := FlightRequest{}
	err := c.Bind(&req)

	if err != nil {
		resp := Response{
			Msg: err.Error(),
		}
		return ResponseUnprocessableEntity(c, resp)
	}
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		resp := Response{
			Msg: "invalid time format",
		}
		return ResponseUnprocessableEntity(c, resp)
	}

	flightReq := &domain.FlightSearchRequest{
		Origin:      req.Origin,
		Destination: req.Destination,
		Date:        date,
	}
	flights, err := f.ctrl.FindBestFlightMatch(c.Request().Context(), flightReq)
	if err != nil {
		resp := Response{
			Msg: err.Error(),
		}
		return ResponseInternalError(c, resp)
	}

	resp := Response{
		Data: flights,
	}

	return ResponseOK(c, resp)
}
