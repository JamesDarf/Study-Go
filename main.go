package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// 템플릿 렌더러 정의
type TemplateRenderer struct {
	templates *template.Template
}

// Render 메서드 구현
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// 렌더러 등록
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	// GET / : 폼 보여주기
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "form.html", nil)
	}) // index.html

	// POST /add : 더하기 처리
	e.POST("/add", func(c echo.Context) error {
		aStr := c.FormValue("a")
		bStr := c.FormValue("b")

		a, _ := strconv.Atoi(aStr)
		b, _ := strconv.Atoi(bStr)
		// Atoi = ASCII to Integer
		// strconv.Atoi() : 문자열을 정수로 변환

		result := a + b

		return c.Render(http.StatusOK, "form.html", map[string]interface{}{
			"Result": result,
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
