package seeders

import (
	"backend/infrastructure"
	"backend/repositories"
	"backend/services"
	"gorm.io/gorm"
	"log"
)

type RateSeeder struct {
	db *gorm.DB
}

// NewRateSeeder initializes and returns a new instance of RateSeeder.
//
// Parameters:
// - db: a pointer to a gorm.DB object representing the database connection.
//
// Returns:
// - a pointer to a RateSeeder object.
func NewRateSeeder(db *gorm.DB) *RateSeeder {
	return &RateSeeder{db: db}
}

// Seed fetches rates from an API, saves them using the rate service, and logs an error if saving fails.
//
// It does not take any parameters.
// No return value.
func (s *RateSeeder) Seed() {
	rateRepository := repositories.NewRateGormRepository(infrastructure.GetDBConnection())
	rateService := services.NewRateService(rateRepository)

	rates := services.GetAllRatesFromAPI()

	err := rateService.SaveAll(rates)

	if err != nil {
		log.Printf("[error] Failed to save rates to database: %v", err)
	}
}
