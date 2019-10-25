package usecase

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
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
