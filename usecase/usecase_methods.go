package usecase

import (
	"github.com/ValeryBMSTU/Go_Tarantool_Task/repository"
	"sync"
)

func (use *Usecase) NewUsecase(mu *sync.Mutex, rep repository.IRepository) {
	use.Repository = rep
	use.Mu = mu
}