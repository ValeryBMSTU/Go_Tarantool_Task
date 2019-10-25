package Go_Tarantool_Task

import (
	"github.com/ValeryBMSTU/Go_Tarantool_Task/delivery"
	"github.com/ValeryBMSTU/Go_Tarantool_Task/repository"
	"github.com/ValeryBMSTU/Go_Tarantool_Task/usecase"
	"github.com/labstack/echo"
	"sync"
)

func main() {
	e := echo.New()

	handlers := delivery.Handlers{}
	var mutex sync.Mutex
	rep := repository.Repository{}

	useCase := usecase.Usecase{}
	useCase.NewUseCase(&mutex, &rep)
	handlers.NewHandlers(e, &useCase)
	e.Logger.Warnf("start listening on %s", 127.0.0.1:8080")
	err = e.Start("127.0.0.1:8080")
	if err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("shutdown")

}
