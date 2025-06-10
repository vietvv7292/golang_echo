package main

import (
    "html/template"
    "io"
    "github.com/labstack/echo/v4"
    "app/router"
)

// Custom template renderer
type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// Cấu hình template
    t := &Template{
        templates: template.Must(template.ParseGlob("views/*.html")),
    }
    e.Renderer = t


	// cấu hình router
	router.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
