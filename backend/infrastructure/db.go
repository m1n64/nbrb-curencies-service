package infrastructure

import (
	"backend/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

var dbConnect *gorm.DB

// InitDBConnection initializes the database connection.
//
// It retrieves the necessary environment variables for the database connection: DB_USER, DB_PASS, DB_NAME, DB_HOST, and DB_PORT.
// Then it creates a DSN (Data Source Name) string using the retrieved environment variables.
// Next, it opens a connection to the database using the gorm package and the created DSN string.
// If there is an error during the connection process, it panics with the message "failed to connect to database".
func InitDBConnection() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	for i := 0; i < 5; i++ {
		dbConnect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		fmt.Printf("Failed to connect to MySQL, retrying... (attempt %d)\n", i+1)
		time.Sleep(5 * time.Second)
	}
}

// GetDBConnection returns the database connection.
//
// No parameters.
// Returns a pointer to a gorm.DB object.
func GetDBConnection() *gorm.DB {
	return dbConnect
}

// MigrateDatabase migrates the database schema to include the `models.Rate` table.
//
// This function connects to the database using the `dbConnect` connection object,
// and calls the `AutoMigrate` method to automatically update the schema with the
// necessary columns and constraints for the `models.Rate` struct.
//
// No parameters are required.
// No return values.
func MigrateDatabase() {
	GetDBConnection().AutoMigrate(&models.Currency{}, &models.Rate{})
}
