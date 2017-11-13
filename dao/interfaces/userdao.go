package interfaces

import "../../models"

type UserDAO interface {
	Create(u models.User) error
	GetByID(id int) (models.User, error)
	GetAll() ([]models.User, error)
	Update(u models.User) error
	Delete(i int) error
}
