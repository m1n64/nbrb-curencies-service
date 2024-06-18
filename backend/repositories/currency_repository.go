package repositories

import "backend/models"

type CurrencyRepository interface {
	Create(currency *models.Currency) error
	SaveMany(currencies []models.Currency) error
	Update(currency *models.Currency) error
	Delete(id uint) error
	FindByID(id uint) (*models.Currency, error)
	FindAll() ([]models.Currency, error)
}
