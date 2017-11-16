package interfaces

import "../../models"

type EntradaDAO interface {
	Create(e *models.Entrada) error
	GetAll() ([]models.Entrada, error)
	GetByID(id int) (models.Entrada, error)
	GetOne(e models.Entrada) (models.Entrada, error)
	Update(e models.Entrada) (models.Entrada, error)
	Delete(id int) error
}
