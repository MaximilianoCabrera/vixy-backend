package interfaces

import "../../models"

type ComentarioEntradaDAO interface {
	GetAll(idComentario int, idEntrada int) ([]models.ComentarioEntrada, error)
	GetOne(ce models.ComentarioEntrada) (models.ComentarioEntrada, error)
	Update(ce models.ComentarioEntrada) error
	Delete(idComentario int, idEntrada int) error
}

