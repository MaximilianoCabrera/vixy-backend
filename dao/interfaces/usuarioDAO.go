package interfaces

import "../../models"

type UsuarioDAO interface {
	Create(u *models.Usuario, ) (string, error)
	GetAll() ([]models.Usuario, error)
	GetByID(id int) (models.Usuario, error)
	GetOne(u models.Usuario) (models.Usuario, error)
	Update(u models.Usuario) (models.Usuario, error)
	Delete(id int) error
}
