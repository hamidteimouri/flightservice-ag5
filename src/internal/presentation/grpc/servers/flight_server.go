package servers

import (
	"ag5/internal/domain"
	"ag5/internal/presentation/grpc/pbs"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type FlightServer struct {
	ctrl *domain.FlightController
}

func NewFlightServer(ctrl *domain.FlightController) *FlightServer {
	return &FlightServer{ctrl: ctrl}
}

func (f FlightServer) BestMatch(ctx context.Context, request *pbs.BestMatchRequest) (*pbs.BestMatchReply, error) {
	logrus.WithFields(logrus.Fields{
		"request": request,
	}).Trace()
	date, err := time.Parse("2006-01-02", request.GetDate())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid date format").Err()
	}
	flightRequest := &domain.FlightSearchRequest{
		Origin:      request.GetOrigin(),
		Destination: request.GetDestination(),
		Date:        date,
	}
	flights, err := f.ctrl.FindBestFlightMatch(ctx, flightRequest)

	if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	results := make([]*pbs.Flight, len(flights))

	for i, flight := range flights {
		results[i] = &pbs.Flight{
			FlightId:           flight.FlightId,
			FlightNo:           flight.FlightNo,
			ScheduledDeparture: timestamppb.New(flight.ScheduledDeparture),
			ScheduledArrival:   timestamppb.New(flight.ScheduledArrival),
			ArrivalAirport:     flight.ArrivalAirport,
			DepartureAirport:   flight.DepartureAirport,
			Status:             flight.Status,
			Date:               flight.ScheduledDeparture.Format("2006-01-02"), // this is just for human test!
		}
	}

	return &pbs.BestMatchReply{
		Flights: results,
	}, nil
}
