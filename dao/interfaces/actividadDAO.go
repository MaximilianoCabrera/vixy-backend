package interfaces

import "../../models"

type ActividadDAO interface {
	Create(a models.Actividad) (string, error)
	GetAll() ([]models.Actividad, error)
	GetByID(id int) (models.Actividad, error)
	GetOne(a models.Actividad) (models.Actividad, error)
	Update(a models.Actividad) error
	Delete(id int) error
}