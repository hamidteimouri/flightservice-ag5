package routes

import (
	"ag5/cmd/dependancy_injection"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	/* handlers */
	flightHandler := dependancy_injection.FlightHandler()

	// initiate a route group
	g := e.Group("api")
	g.GET("/flights", flightHandler.BestMatch)
}
