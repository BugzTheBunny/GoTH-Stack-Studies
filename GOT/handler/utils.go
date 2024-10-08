package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func renderTemplate(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
