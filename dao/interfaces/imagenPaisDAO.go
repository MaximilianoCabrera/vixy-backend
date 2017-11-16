package interfaces

import "../../models"

type ImagenPaisDAO interface {
	GetAll(idPais int, idImagen int) ([]models.ImagenPais, error)
	GetOne(ip models.ImagenPais) (models.ImagenPais, error)
	Update(ip models.ImagenPais) error
	Delete(idPais int, idCiudad int) error
}
