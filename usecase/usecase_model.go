package usecase

import (
	"github.com/ValeryBMSTU/Go_Tarantool_Task/models"
	"github.com/ValeryBMSTU/Go_Tarantool_Task/repository"
	"sync"
)

type Usecase struct {
	Repository repository.IRepository
	Mu          *sync.Mutex
}

type IUsecase interface {
	AddKeyValue(KeyValue models.PostKeyValue) (models.OutKeyValue, error)
	GetValue(Key string) (models.OutKeyValue, error)
	Delete(Key string) (models.OutKeyValue, error)
	Set(KeyValue models.PostKeyValue) (models.OutKeyValue, error)
}
