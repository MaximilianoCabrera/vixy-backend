package interfaces

import "../../models"

type ContinenteDAO interface {
	Create(c models.Continente) error
	GetAll() ([]models.Continente, error)
	GetByID(id int) (models.Continente, error)
	GetOne(c models.Continente) (models.Continente, error)
	Update(c models.Continente) error
	Delete(id int) error
}