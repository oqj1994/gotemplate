package app

import (
	"fmt"
	"log"
	"net/http"
	"projectTemplate/modle"
	"projectTemplate/public"
	template "projectTemplate/templ"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	r       *echo.Echo
	Address string
}

func (a *App) Start() error {
	fmt.Println("app start ,run on port: ", a.Address)
	a.r.Use(middleware.Logger())
	return http.ListenAndServe(a.Address, a.r)
}

func (a *App) Register() error {
	a.r.GET("/", func(c echo.Context) error {
		boxs := modle.Boxs
		r := template.Index("Just some box", boxs)

		return r.Render(c.Request().Context(), c.Response())

	})
	a.r.GET("/public/*", func(c echo.Context) error {
		url := c.Request().URL.Path
		path := strings.TrimPrefix(url, "/")
		fmt.Println("path ......", path)
		b, err := public.FS.ReadFile(strings.Split(path, "/")[1])
		if err != nil {
			log.Println(err)
		}
		_, err = c.Response().Write(b)
		return err
	})
	a.r.POST("/box", func(c echo.Context) error {
		title := c.FormValue("title")
		modle.Boxs = append(modle.Boxs, modle.Box{
			UUID:  uuid.New(),
			Title: title,
			State: 1,
			Count: 24,
		})
		return c.Redirect(http.StatusFound, "/")
	})
	return nil
}

func NewApp(address string) *App {
	return &App{
		r:       echo.New(),
		Address: address,
	}
}
