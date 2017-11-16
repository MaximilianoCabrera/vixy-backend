package interfaces

import "../../models"

type ImagenDAO interface {
	Create(i *models.Imagen) (models.Imagen, error)
	GetAll() ([]models.Imagen, error)
	GetByID(id int) (models.Imagen, error)
	GetOne(i models.Imagen) (models.Imagen, error)
	Update(i models.Imagen) error
	Delete(id int) error
}