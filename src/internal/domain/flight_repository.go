package domain

import (
	"context"
)

type FlightRepository interface {
	FindBestMatches(ctx context.Context, request *FlightSearchRequest) ([]Flight, error)
}
