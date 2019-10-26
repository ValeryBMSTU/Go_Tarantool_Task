package middlewares

import (
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	tarantool "github.com/tarantool/go-tarantool"
)

func PanicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		defer func() {
			if panicErr := recover(); panicErr != nil {
				ctx.Logger().Error("recovered ", panicErr)
				err = &echo.HTTPError{Code: 500, Message: "Internal server error"}
			}
		}()
		err = next(ctx)
		return err
	}
}

func CustomErrorHandler(err error, ctx echo.Context) {
	var jsonError error
	switch err := errors.Cause(err); err.(type) {
	case *echo.HTTPError:
		ctx.Logger().Warn(err)
		jsonError = ctx.JSON(err.(*echo.HTTPError).Code, struct {
			Info string `json:"info"`
		}{Info: err.(*echo.HTTPError).Message.(string)})
	case tarantool.Error:
		if err.(tarantool.Error).Code == 3 {
			ctx.Logger().Warn(err)
			jsonError = ctx.JSON(409, struct {
				Info string `json:"info"`
			}{Info: err.Error()})
			return
		}
		if err.(tarantool.Error).Code == 4 {
			ctx.Logger().Warn(err)
			jsonError = ctx.JSON(404, struct {
				Info string `json:"info"`
			}{Info: err.Error()})
			return
		}
		return
	case nil:
		return
	default:
		ctx.Logger().Info(err)
		jsonError = ctx.JSON(400, struct {
			Info string `json:"info"`
		}{Info: err.Error()})
	}
	if jsonError != nil {
		ctx.Logger().Error("Server cant repay response")
	}
}
