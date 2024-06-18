package repositories

import (
	"backend/models"
	"time"
)

type RateRepository interface {
	Create(rate *models.Rate) error
	SaveMany(rates []models.Rate) error
	Update(rate *models.Rate) error
	Delete(id uint) error
	FindByID(id uint) (*models.Rate, error)
	FindByDate(date time.Time) ([]models.Rate, error)
	FindAll() ([]models.Rate, error)
}
