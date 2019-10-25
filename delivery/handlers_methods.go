package delivery

import (
	"github.com/ValeryBMSTU/Go_Tarantool_Task/usecase"
	"github.com/labstack/echo"
)

func (h *Handlers) NewHandlers(e *echo.Echo, IUsecase usecase.IUsecase) {
	h.Usecase = IUsecase


	// e.POST("/kv", h.PostKeyValue)
	// e.PUT( "/kv/:id", h.PutKeyValue)
	// e.GET("/kv/:id", h.GetKeyValue)
	// e.DELETE( "/kv/:id", h.DeleteKeyValue)
}

