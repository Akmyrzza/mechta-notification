package repository

import "github.com/akmyrzza/mechta-notification/internal/models"

type Repository interface {
	GetAll() ([]string, error)

	GetAllChannels() ([]models.Channel, error)
}
