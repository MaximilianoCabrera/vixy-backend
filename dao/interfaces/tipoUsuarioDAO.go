package interfaces

import "../../models"

type TipoUsuarioDAO interface {
	Create(tu models.TipoUsuario) error
	GetByID(id int) (models.TipoUsuario, error)
	GetOne(tu models.TipoUsuario) (models.TipoUsuario, error)
	GetAll() ([]models.TipoUsuario, error)
	Update(tu models.TipoUsuario) error
	Delete(i int) error
}