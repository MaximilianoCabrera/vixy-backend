package interfaces

import "../../models"

type ImagenCiudadDAO interface {
	GetAll(idCiudad int, idImagen int) ([]models.ImagenCiudad, error)
	GetOne(ic models.ImagenCiudad) (models.ImagenCiudad, error)
	Update(ic models.ImagenCiudad) error
	Delete(idCiudad, idImagen int) error
}
