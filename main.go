package main

import (
	_ "errors"
	"github.com/ValeryBMSTU/Go_Tarantool_Task/delivery"
	customMiddlewares "github.com/ValeryBMSTU/Go_Tarantool_Task/middlewares"
	"github.com/ValeryBMSTU/Go_Tarantool_Task/repository"
	"github.com/ValeryBMSTU/Go_Tarantool_Task/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"sync"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "${time_rfc3339}, method = ${method}, uri = ${uri}, status = ${status}, remote_ip = ${remote_ip}\n"}))
	e.Use(customMiddlewares.PanicMiddleware)
	e.HTTPErrorHandler = customMiddlewares.CustomErrorHandler

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
	
	err = e.Start("127.0.0.1:8081")
	if err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("shutdown")
}
