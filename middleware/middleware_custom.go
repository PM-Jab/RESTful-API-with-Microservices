package middleware_custom

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Hello() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		fmt.Println("middle 1")
		return func(c echo.Context) (err error) {
			fmt.Println("middle 2")
			return next(c)
		}
	}
}
