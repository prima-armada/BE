package handlecontract

import "github.com/labstack/echo/v4"

type HandleUser interface {
	Register(e echo.Context) error
}
