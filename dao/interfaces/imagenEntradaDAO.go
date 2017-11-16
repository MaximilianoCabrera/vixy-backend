package interfaces

import "../../models"

type ImagenEntradaDAO interface {
	GetAll(idEntrada int, idImagen int) ([]models.ImagenEntrada, error)
	GetOne(ie models.ImagenEntrada) (models.ImagenEntrada, error)
	Update(ie models.ImagenEntrada) error
	Delete(idEntrada, idImagen int) error
}