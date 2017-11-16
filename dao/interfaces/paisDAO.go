package interfaces

import "../../models"

type PaisDAO interface {
	Create(p models.Pais) error
	GetAll() ([]models.Pais, error)
	GetByID(id int) (models.Pais, error)
	GetOne(p models.Pais) (models.Pais, error)
	Update(p models.Pais) error
	Delete(id int) error
}
