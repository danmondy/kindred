package main

import (
	"fmt"

	//templates
	"github.com/a-h/templ"

	"github.com/danmondy/kindred/models"
	"github.com/danmondy/kindred/templates"

	"github.com/labstack/echo/v4"
)

var cfg *models.Config

const CONFIG_FILE string = "config.json"

func main() {
	e := echo.New()

	e.Static("/assets", cfg.StaticFolder)

	e.GET("/", IndexHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}

func init() {
	cfg = &models.Config{}

	err := models.MapJson(CONFIG_FILE, cfg)
	if err != nil {
		fmt.Println(err)
	}
}

// this is used by ever handler to render a view with Echo and Templ
func render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func IndexHandler(c echo.Context) error {
	return render(c, 200, templates.IndexPage())
}
