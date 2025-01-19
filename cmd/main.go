package main

import (
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	tmpl *template.Template
}

func newTemplate() *Template {
	return &Template{
		tmpl: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}

type Data struct {
	Title string
}

func newData() *Data {
	return &Data{
		Title: "Hello World",
	}
}

func main() {
	e := echo.New()
	d := newData()

	e.Renderer = newTemplate()
	e.Use(middleware.Logger())

	// Serve static assets correctly
	e.Static("/css", "static/css")       // Serve CSS files under /css
	e.Static("/assets", "static/assets") // Serve other assets under /assets

	// Route handler for the home page
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index.html", d)
	})

	// Start the server on port 5000
	log.Println("Server started on http://localhost:5000")
	e.Logger.Fatal(e.Start(":5000"))
}
