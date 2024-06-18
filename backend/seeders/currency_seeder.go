package seeders

import (
	"backend/infrastructure"
	"backend/repositories"
	"backend/services"
	"gorm.io/gorm"
	"log"
)

type CurrencySeeder struct {
	db *gorm.DB
}

// NewCurrencySeeder initializes and returns a new instance of CurrencySeeder.
//
// Parameters:
// - db: a pointer to a gorm.DB object representing the database connection.
//
// Returns:
// - a pointer to a CurrencySeeder object.
func NewCurrencySeeder(db *gorm.DB) *CurrencySeeder {
	return &CurrencySeeder{db: db}
}

// Seed seeds the currency repository with currencies obtained from the API.
//
// It does not take any parameters.
// It does not return anything.
func (s *CurrencySeeder) Seed() {
	currencyRepository := repositories.NewCurrencyGormRepository(infrastructure.GetDBConnection())

	currencies := services.GetAllCurrenciesFromAPI()

	err := currencyRepository.SaveMany(currencies)
	if err != nil {
		log.Fatalf("Error seeding currencies: %v", err)
	}
}
