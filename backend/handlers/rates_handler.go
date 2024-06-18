package handlers

import (
	"backend/infrastructure"
	"backend/models"
	"backend/repositories"
	"backend/services"
	"github.com/araddon/dateparse"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AllCurrencyRates retrieves all currency rates using the provided rate service.
//
// Parameters:
// - c: The Gin context for handling the HTTP request.
// Return type(s).
func AllCurrencyRates(c *gin.Context) {
	filterCurrencyRates(c, func(rateService *services.RateService) ([]models.Rate, error) {
		return rateService.All()
	})
}

// CurrencyRateByDate retrieves currency rates by date.
//
// This function takes a Gin context as a parameter and retrieves the date string from the context's parameters.
// It then parses the date string into a time.Time object using the dateparse.ParseAny function.
// If the parsing fails, it returns a JSON response with a bad request error.
// Otherwise, it calls the filterCurrencyRates function with the parsed date and a callback function that retrieves the rates from the rate service.
// The callback function takes a rate service as a parameter and returns a slice of models.Rate and an error.
// The function returns nothing.
func CurrencyRateByDate(c *gin.Context) {
	dateString := c.Param("date")

	date, err := dateparse.ParseAny(dateString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format", "success": false})
		return
	}

	filterCurrencyRates(c, func(rateService *services.RateService) ([]models.Rate, error) {
		return rateService.FindByDate(date)
	})
}

// filterCurrencyRates handles currency rate filtering based on the provided callback function.
//
// Parameters:
// - c: The Gin context for handling the HTTP request.
// - callback: A function that takes a RateService and returns a slice of Rate models and an error.
// Return type(s).
func filterCurrencyRates(c *gin.Context, callback func(rateService *services.RateService) ([]models.Rate, error)) {
	rateRepository := repositories.NewRateGormRepository(infrastructure.GetDBConnection())
	rateService := services.NewRateService(rateRepository)

	rates, err := callback(rateService)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rates,
	})
}
