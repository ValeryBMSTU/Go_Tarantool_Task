package delivery

import (
	"github.com/ValeryBMSTU/Go_Tarantool_Task/usecase"
	"github.com/labstack/echo"
)

func (h *Handlers)NewHandlers(e *echo.Echo, IUsecase usecase.IUsecase) {
	h.Usecase = IUsecase


}
