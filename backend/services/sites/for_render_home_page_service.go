package sites

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func RenderHomePageService(c *gin.Context) {
	data := struct {
		Title string
	}{
		Title: "Página de Busqueda",
	}

	// Parsear el archivo de template
	tmpl, err := template.ParseFiles("../frontend/static/templates/home.html")

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Renderizar el template con los datos
	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
}
