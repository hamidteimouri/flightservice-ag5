package dependancy_injection

import (
	"ag5/internal/data/postgres"
	"ag5/internal/domain"
	"ag5/internal/presentation/grpc/servers"
	"ag5/internal/presentation/http/handlers"
	"fmt"
	"github.com/hamidteimouri/gommon/htenvier"
	"github.com/sirupsen/logrus"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB

	/* Controllers variable */
	flightController *domain.FlightController
	dbDatasource     *postgres.Postgres

	/* Handlers variables */
	flightHandler *handlers.FlightHandler

	/* GRPC variables */
	flightServer *servers.FlightServer
)

func DB() *gorm.DB {
	if db != nil {
		return db
	}

	var err error
	dbHost := htenvier.Env("DB_HOST")
	dbPort := htenvier.Env("DB_PORT")
	dbName := htenvier.Env("DB_NAME")
	dbUsername := htenvier.Env("DB_USERNAME")
	dbPassword := htenvier.Env("DB_PASSWORD")
	dbTimezone := htenvier.Env("DB_TIMEZONE")

	// logger of gorm
	gormLogger := logger.Default.LogMode(logger.Silent)

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v",
		dbHost, dbUsername, dbPassword, dbName, dbPort, dbTimezone)
	db, err = gorm.Open(gormPostgres.Open(dsn), &gorm.Config{Logger: gormLogger})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("database connection error")
	}

	return db
}

/*********** Datasource ***********/

func DbDatasource() *postgres.Postgres {
	if dbDatasource != nil {
		return dbDatasource
	}
	dbDatasource = postgres.NewPostgres(db)
	return dbDatasource
}

/*********** Domain ***********/

func FlightDomain() *domain.FlightController {
	if flightController != nil {
		return flightController
	}
	flightController = domain.NewFlightController(DbDatasource())
	return flightController
}

/*********** Handlers ***********/

func FlightHandler() *handlers.FlightHandler {
	if flightHandler != nil {
		return flightHandler
	}
	flightHandler = handlers.NewFlightHandler(FlightDomain())
	return flightHandler
}

/*********** GRPC ***********/

func GrpcFlightServer() *servers.FlightServer {
	if flightServer != nil {
		return flightServer
	}
	flightServer = servers.NewFlightServer(FlightDomain())
	return flightServer
}
