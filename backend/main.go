package main

import (
	"backend/handlers"
	"backend/infrastructure"
	"backend/seeders"
	"backend/workers"
	"github.com/gin-gonic/gin"
)

func main() {
	infrastructure.InitDBConnection()
	infrastructure.MigrateDatabase()
	SeedDatabase()

	r := gin.Default()
	routes(r)

	go workers.SaveCurrencyRates()

	r.Run(":80")
}

func routes(r *gin.Engine) {
	r.GET("/ping", handlers.PingHandler)
	r.GET("/rates", handlers.AllCurrencyRates)
	r.GET("/rates/:date", handlers.CurrencyRateByDate)
	r.GET("/swagger/docs.json", func(c *gin.Context) {
		c.File("swagger/docs.json")
	})
}

// SeedDatabase seeds the database with currency and rate data.
//
// This function creates a new instance of CurrencySeeder and RateSeeder
// using the GetDBConnection function to establish a database connection.
// It then calls the Seed method of each seeder to populate the database
// with currency and rate data.
func SeedDatabase() {
	currencySeeder := seeders.NewCurrencySeeder(infrastructure.GetDBConnection())
	currencySeeder.Seed()

	rateSeeder := seeders.NewRateSeeder(infrastructure.GetDBConnection())
	rateSeeder.Seed()
}
