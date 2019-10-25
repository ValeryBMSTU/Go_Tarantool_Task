package main

import (
	_ "errors"
	"github.com/ValeryBMSTU/Go_Tarantool_Task/repository"
	"log"
	"github.com/ValeryBMSTU/Go_Tarantool_Task/delivery"
	"github.com/ValeryBMSTU/Go_Tarantool_Task/usecase"
	"github.com/labstack/echo"
	"sync"
)

func main() {
	e := echo.New()

	rep := repository.Repository{}
	err := rep.NewRepository()
	if err != nil {
		log.Fatalf("cant connect to tarantool: %s\n", err)
	}

	uc := usecase.Usecase{}
	uc.NewUsecase(&sync.Mutex{}, &rep)

	handlers := delivery.Handlers{}
	handlers.NewHandlers(e, &uc)

	e.Logger.Warnf("start listening on %s", "127.0.0.1:8080")
	
	err = e.Start("127.0.0.1:8080")
	if err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("shutdown")
}
