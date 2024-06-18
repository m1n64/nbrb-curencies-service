package repositories

import (
	"backend/models"
	"gorm.io/gorm"
)

type CurrencyGormRepository struct {
	db *gorm.DB
}

// NewCurrencyGormRepository initializes a new RateGormRepository with the provided *gorm.DB.
//
// Parameters:
// - db: *gorm.DB object for database operations.
// Returns a pointer to a RateGormRepository.
func NewCurrencyGormRepository(db *gorm.DB) *CurrencyGormRepository {
	return &CurrencyGormRepository{
		db: db,
	}
}

// Create creates a new currency rate in the database.
//
// rate: a pointer to a models.Currency struct
// error: an error if any
func (r *CurrencyGormRepository) Create(rate *models.Currency) error {
	return r.db.Create(rate).Error
}

// SaveMany saves multiple currency rates to the database using GORM.
//
// It takes a slice of models.Currency as a parameter and returns an error.
// The function starts a transaction and iterates over the rates, saving each one to the database.
// If any rate fails to be saved, the transaction is rolled back and the error is returned.
// If all rates are successfully saved, the transaction is committed and no error is returned.
func (r *CurrencyGormRepository) SaveMany(rates []models.Currency) error {
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

// Update updates a currency in the database using the Gorm repository.
//
// Parameters:
// - rate: A pointer to a models.Currency object representing the currency to be updated.
//
// Returns:
// - error: An error object if there was an error updating the currency, otherwise nil.
func (r *CurrencyGormRepository) Update(rate *models.Currency) error {
	return r.db.Save(rate).Error
}

// Delete deletes a currency from the database using the Gorm repository.
//
// Parameters:
// - id (uint): The ID of the currency to delete.
//
// Returns:
// - error: An error object if there was an error deleting the currency, otherwise nil.
func (r *CurrencyGormRepository) Delete(id uint) error {
	return r.db.Delete(&models.Currency{}, id).Error
}

// FindByID retrieves a currency by its ID from the database using GORM.
//
// Parameters:
// - id (uint): The ID of the currency to retrieve.
//
// Returns:
// - *models.Currency: A pointer to the retrieved currency.
// - error: An error object if the retrieval operation fails.
func (r *CurrencyGormRepository) FindByID(id uint) (*models.Currency, error) {
	var rate models.Currency
	err := r.db.First(&rate, id).Error

	return &rate, err
}

// FindAll retrieves all rates from the database.
//
// Returns:
// - []models.Currency: A slice of all the retrieved rates.
// - error: An error object if the retrieval operation fails.
func (r *CurrencyGormRepository) FindAll() ([]models.Currency, error) {
	var rates []models.Currency
	err := r.db.Find(&rates).Error

	return rates, err
}
