package main

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/cgradwohl/echo-templ-2/view"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func main() {

	app := echo.New()

	app.Start(":80")
	fmt.Println("hello creature ...")

	app.GET("/", func(c echo.Context) error {
		return render(c, view.Hello("chris"))
	})
}
