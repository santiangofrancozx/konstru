package sites

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderRegisterTemplateService(c *gin.Context) {
	// Definir los datos a pasar al template (si es necesario)
	data := struct {
		Title string
	}{
		Title: "Página de Inicio de Sesión",
	}

	// Parsear el archivo de template
	tmpl, err := template.ParseFiles("../frontend/static/templates/registro.html")

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
