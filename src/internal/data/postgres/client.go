package postgres

import (
	"ag5/internal/domain"
	"context"
	"errors"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgres(db *gorm.DB) *Postgres {
	return &Postgres{db: db}
}

func (m *Postgres) FindBestMatches(ctx context.Context, request *domain.FlightSearchRequest) ([]domain.Flight, error) {
	var flightsModel []FlightModel

	result := m.db.Model(&FlightModel{}).
		Where("status = ?", "Scheduled").
		Where("CAST(scheduled_departure AS DATE) = ?", request.Date.Format("2006-01-02")).
		Where("departure_airport = ?", request.Origin).
		Where("arrival_airport = ?", request.Destination).
		//Where("scheduled_departure >= ? ", request.Date).
		Order("scheduled_departure desc").Find(&flightsModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	flightsEntity := make([]domain.Flight, len(flightsModel))
	for i, model := range flightsModel {
		flightsEntity[i] = model.ConvertModelToEntity()
	}

	return flightsEntity, nil
}
