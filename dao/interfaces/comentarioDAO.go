package interfaces

import "../../models"

type ComentarioDAO interface {
	Create(C models.Comentario) (string, error)
	GetAll() ([]models.Comentario, error)
	GetByID(id int) (models.Comentario, error)
	GetOne(C models.Comentario) (models.Comentario, error)
	Update(C models.Comentario) error
	Delete(id int) error
}