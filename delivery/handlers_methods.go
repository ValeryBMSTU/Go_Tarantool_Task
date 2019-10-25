package delivery

import (
	"encoding/json"
	"github.com/ValeryBMSTU/Go_Tarantool_Task/models"
	"github.com/ValeryBMSTU/Go_Tarantool_Task/usecase"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func (h *Handlers) NewHandlers(e *echo.Echo, IUsecase usecase.IUsecase) {
	h.Usecase = IUsecase


	e.POST("/kv", h.PostKeyValue)
	e.PUT( "/kv/:id", h.PutKeyValue)
	e.GET("/kv/:id", h.GetKeyValue)
	e.DELETE( "/kv/:id", h.DeleteKeyValue)
}

func (h *Handlers) PostKeyValue (ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	decoder := json.NewDecoder(ctx.Request().Body)
	newKeyValue := new(models.PostKeyValue)
	if err := decoder.Decode(newKeyValue); err != nil {
		return err
	}

	res, err := h.Usecase.AddKeyValue(*newKeyValue)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(ctx.Response())
	if err := encoder.Encode(res); err != nil {
		return err
	}

	return nil
}

func (h *Handlers) GetKeyValue (ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	key := ctx.Param("id")

	keyValue, err := h.Usecase.GetValue(key)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(ctx.Response())
	if err := encoder.Encode(keyValue); err != nil {
		return err
	}

	return nil
}

func (h *Handlers) DeleteKeyValue (ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	key := ctx.Param("id")

	res, err := h.Usecase.Delete(key)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(ctx.Response())
	if err := encoder.Encode(res); err != nil {
		return err
	}

	return nil
}

func (h *Handlers) PutKeyValue (ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	key := ctx.Param("id")

	decoder := json.NewDecoder(ctx.Request().Body)
	newKeyValue := new(models.PostKeyValue)
	newPutValue := new(models.PutValue)
	if err := decoder.Decode(newPutValue); err != nil {
		return err
	}
	newKeyValue.Key = key
	newKeyValue.Value = newPutValue.Value

	res, err := h.Usecase.Set(*newKeyValue)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(ctx.Response())
	if err := encoder.Encode(res); err != nil {
		return err
	}

	return nil
}