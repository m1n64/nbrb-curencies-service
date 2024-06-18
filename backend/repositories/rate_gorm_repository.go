package repositories

import (
	"backend/models"
	"gorm.io/gorm"
	"time"
)

type RateGormRepository struct {
	db *gorm.DB
}

// NewRateGormRepository initializes a new RateGormRepository with the provided *gorm.DB.
//
// Parameters:
// - db: *gorm.DB object for database operations.
// Returns a pointer to a RateGormRepository.
func NewRateGormRepository(db *gorm.DB) *RateGormRepository {
	return &RateGormRepository{
		db: db,
	}
}

// Create saves a new rate to the database using GORM.
//
// It takes a pointer to a models.Rate as a parameter and returns an error.
func (r *RateGormRepository) Create(rate *models.Rate) error {
	return r.db.Create(rate).Error
}

// SaveMany saves multiple rates to the database using GORM.
//
// It takes a slice of pointers to models.Rate as a parameter and returns an error.
// The function starts a transaction and iterates over the rates, saving each one to the database.
// If any rate fails to be saved, the transaction is rolled back and the error is returned.
// If all rates are successfully saved, the transaction is committed and no error is returned.
func (r *RateGormRepository) SaveMany(rates []models.Rate) error {
	tx := r.db.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	for _, rate := range rates {
		if err := tx.Save(&rate).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// Update updates the given Rate in the database using the Gorm repository.
//
// Parameters:
// - rate: A pointer to a Rate model object that represents the rate to be updated.
//
// Returns:
// - error: An error object if there was an error updating the rate, otherwise nil.
func (r *RateGormRepository) Update(rate *models.Rate) error {
	return r.db.Save(rate).Error
}

// Delete deletes a rate from the database using the Gorm repository.
//
// Parameters:
// - id: The ID of the rate to be deleted.
//
// Returns:
// - error: An error object if there was an error deleting the rate, otherwise nil.
func (r *RateGormRepository) Delete(id uint) error {
	return r.db.Delete(&models.Rate{}, id).Error
}

// FindByID retrieves a rate from the database based on the provided ID.
//
// Parameters:
// - id: The ID of the rate to retrieve.
// Returns:
// - *models.Rate: A pointer to the retrieved rate.
// - error: An error object if the retrieval operation fails.
func (r *RateGormRepository) FindByID(id uint) (*models.Rate, error) {
	var rate models.Rate
	err := r.db.First(&rate, id).Error

	return &rate, err
}

// FindAll retrieves all rates from the database.
//
// Returns:
// - []models.Rate: A slice of Rate objects representing all rates in the database.
// - error: An error object if the retrieval operation fails.
func (r *RateGormRepository) FindAll() ([]models.Rate, error) {
	var rates []models.Rate
	err := r.db.Preload("Currency").Find(&rates).Error

	return rates, err
}

// FindByDate retrieves a rate from the database based on the provided date.
//
// Parameters:
// - date: The date of the rate to retrieve.
// Returns:
// - *models.Rate: A pointer to the retrieved rate.
// - error: An error object if the retrieval operation fails.
func (r *RateGormRepository) FindByDate(date time.Time) ([]models.Rate, error) {
	var rates []models.Rate
	err := r.db.Preload("Currency").Find(&rates, "date = ?", date).Error

	return rates, err
}
