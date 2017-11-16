package interfaces

import "../../models"

type CiudadDAO interface {
	Create(c models.Ciudad) (string, error)
	GetAll() ([]models.Ciudad, error)
	GetByID(id int) (models.Ciudad, error)
	GetOne(c models.Ciudad) (models.Ciudad, error)
	Update(c models.Ciudad) error
	Delete(id int) error
}