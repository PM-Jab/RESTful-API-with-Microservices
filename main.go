package main

import (
	"net/http"

	mc "github.com/PM-Jab/cookie_shop/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))
	g.GET("/home", func(c echo.Context) error {
		return c.String(http.StatusOK, "Admin's Page")
	})

	a := e.Group("/cookie")
	a.Use(mc.Hello())
	a.GET("/home", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	a.GET("/menu", func(c echo.Context) error {
		return c.String(http.StatusOK, "Sweet cookie")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
