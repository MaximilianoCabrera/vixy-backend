package interfaces


import "../../models"

type GlobalDAO interface {
	Create(x *models.GlobalModel, model string) (models.GlobalModel, error)
	GetAll(model string) (models.GlobalModels, error)
	GetOne(x models.GlobalModel, model string) (models.GlobalModel, error)
	GetByID(id int, model string) (models.GlobalModel, error)
	Update (x models.GlobalModel, model string) (models.GlobalModel, error)
	Delete(x *models.GlobalModel, model string) (string, error)
}
