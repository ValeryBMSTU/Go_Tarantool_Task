package usecase

import (
	"encoding/json"
	"github.com/ValeryBMSTU/Go_Tarantool_Task/repository"
	"github.com/ValeryBMSTU/Go_Tarantool_Task/pkg/models"
	"io"
	"net/http"
	"sync"
)

type Usecase struct {
	PRepository repository.IRepository
	Mu          *sync.Mutex
}

type IUsecase interface {

}
