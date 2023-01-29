package domain

import (
	"context"
)

type FlightController struct {
	repo FlightRepository
}

func NewFlightController(repo FlightRepository) *FlightController {
	return &FlightController{repo: repo}
}

func (f *FlightController) FindBestFlightMatch(ctx context.Context, request *FlightSearchRequest) ([]Flight, error) {
	return f.repo.FindBestMatches(ctx, request)
}
