package interfaces

import "../../models"

type TopicDAO interface {
	Create(t models.Topic) (string, error)
	GetAll() ([]models.Topic, error)
	GetByID(id int) (models.Topic, error)
	GetOne(t models.Topic) (models.Topic, error)
	Update(t models.Topic) error
	Delete(id int) error
}
