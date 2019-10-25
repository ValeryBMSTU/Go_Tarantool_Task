package usecase

import (
	"github.com/ValeryBMSTU/Go_Tarantool_Task/repository"
	"sync"
)

type Usecase struct {
	Repository repository.IRepository
	Mu          *sync.Mutex
}

type IUsecase interface {
	
}
