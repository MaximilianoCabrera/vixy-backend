package interfaces

import "../../models"

type UsuarioEntradaDAO interface {
	GetAll(idUsuario int, idEntrada int) ([]models.UsuarioEntrada, error)
	GetOne(ue models.UsuarioEntrada) (models.UsuarioEntrada, error)
	Update(ue models.UsuarioEntrada) error
	Delete(idUsuario int, idEntrada int) error
}
