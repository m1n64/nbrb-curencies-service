package workers

import (
	"backend/infrastructure"
	"backend/seeders"
	"github.com/robfig/cron/v3"
	"log"
)

// SaveCurrencyRates updates the currency rates by fetching rates from the API and saving them to the database.
//
// No parameters.
// No return value.
func SaveCurrencyRates() {
	c := cron.New()

	_, err := c.AddFunc("0 0 * * *", func() {
		rateSeeder := seeders.NewRateSeeder(infrastructure.GetDBConnection())
		rateSeeder.Seed()
	})

	if err != nil {
		log.Printf("[error] Error adding cron job: %v", err)
		return
	}

	c.Start()

	select {}
}
