package interfaces


import "../../models"

type GlobalDAO interface {
	Create(x *models.GlobalModel, model string) (models.GlobalModel, error)
	GetAll(model string) (models.GlobalModels, error)
	GetBy(x models.GlobalModel, model string) (models.GlobalModels, error)
	GetByID(id int, model string) (models.GlobalModel, error)
	Update (x models.GlobalModel, model string) (models.GlobalModel, error)
	Delete(x *models.GlobalModel, model string) (string, error)
}
