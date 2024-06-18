package services

import (
	"backend/models"
	"encoding/json"
)

// GetAllCurrenciesFromAPI retrieves all currencies from the API.
//
// No parameters.
// []models.Currency: A slice of Currency models representing the retrieved currencies.
func GetAllCurrenciesFromAPI() []models.Currency {
	body, err := GetRequest("https://api.nbrb.by/exrates/rates?periodicity=0")

	var rates []models.Currency
	err = json.Unmarshal(body, &rates)
	if err != nil {
		panic(err)
	}

	return rates
}
