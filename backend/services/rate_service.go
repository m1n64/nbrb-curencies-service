package services

import (
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"time"
)

type RateService struct {
	repository repositories.RateRepository
}

// NewRateService creates a new instance of the RateService struct with the provided repository.
//
// Parameters:
// - repository: An instance of the RateRepository interface that will be used for storing and retrieving rate data.
//
// Returns:
// - A pointer to the newly created RateService instance.
func NewRateService(repository repositories.RateRepository) *RateService {
	return &RateService{
		repository: repository,
	}
}

// CreateOrUpdate updates an existing Rate record or creates a new one if it doesn't exist.
//
// Parameters:
// - rate: a pointer to a Rate object containing the data to be updated or created.
//
// Returns:
// - error: an error if the update or creation of the Rate record fails.
func (s *RateService) CreateOrUpdate(rate *models.Rate) error {
	return s.repository.Update(rate)
}

func (s *RateService) SaveAll(rates []models.Rate) error {
	return s.repository.SaveMany(rates)
}

func (s *RateService) All() ([]models.Rate, error) {
	return s.repository.FindAll()
}

func (s *RateService) FindByDate(date time.Time) ([]models.Rate, error) {
	return s.repository.FindByDate(date)
}

// GetAllRatesFromAPI retrieves all rates from the NBRB API.
//
// It sends a GET request to the NBRB API endpoint for rates with periodicity=0
// and returns a slice of Rate models representing the retrieved rates.
//
// Returns:
//   - []models.Rate: A slice of Rate models representing the retrieved rates.
//   - None: Panics if there is an error creating the HTTP request or sending it.
//     Panics if there is an error reading the response body or unmarshaling
//     the JSON response into the rates slice.
func GetAllRatesFromAPI() []models.Rate {
	body, err := GetRequest("https://api.nbrb.by/exrates/rates?periodicity=0")

	var rates []models.Rate
	err = json.Unmarshal(body, &rates)
	if err != nil {
		panic(err)
	}

	return rates
}
