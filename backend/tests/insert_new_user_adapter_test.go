package tests

import (
	"awesomeKonstru/backend/config"
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// TestInsertNewUser_Successful testea el caso donde la inserción es exitosa
func TestInsertNewUser_Successful(t *testing.T) {
	config.SetDSN("root:safraroot@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local")

	// Crear usuario de prueba
	user := models.Usuario{
		Nombre:   "John",
		Apellido: "Doe",
		Email:    "john.doe5@example.com",
		Password: "123456",
	}

	// Llamar a la función
	successful, err := Adapters.InsertNewUser(user)

	// Verificaciones
	assert.True(t, successful, "Se esperaba que la inserción fuera exitosa")
	assert.NoError(t, err, "No se esperaba un error")

	// Limpiar datos después de la prueba
	cleanUpInsertedUserData(user.Email)
}

func cleanUpInsertedUserData(email string) {
	// Conectar a la base de datos
	db, err := Connection_Migrates.Connect()
	if err != nil {
		log.Fatalf("Error conectando a la base de datos para limpieza: %v", err)
	}
	defer Connection_Migrates.Disconnect(db)

	// Eliminar el usuario insertado
	err = Queries.QueryDeleteUserByEmail(db, email)
	if err != nil {
		log.Fatalf("Error eliminando usuario insertado: %v", err)
	}
}

func TestInsertNewProjectAdapter_Successful(t *testing.T) {
	// Configuración de prueba
	config.SetDSN("root:safraroot@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local")

	// Crear proyecto de prueba
	project := models.Proyectos{
		Name:        "Proyecto de Ejemplo",
		Descripcion: "Descripción del proyecto",
		Valor:       100000.00,
		TipoObra:    "Construcción",
		UsuarioID:   "usuario_prueba_id", // Asegúrate de que este usuario exista en la base de datos
	}

	// Llamar al adaptador
	err, insertedProject := Adapters.InsertNewProjectAdapter(project)

	// Verificaciones
	assert.NoError(t, err, "No se esperaba un error al insertar el proyecto")
	assert.NotNil(t, insertedProject, "Se esperaba que el proyecto insertado no fuera nulo")
	assert.NotEmpty(t, insertedProject.IDProyecto, "Se esperaba que el ID del proyecto insertado no estuviera vacío")
	assert.Equal(t, project.Name, insertedProject.Name, "El nombre del proyecto insertado no coincide con el esperado")
	assert.Equal(t, project.Descripcion, insertedProject.Descripcion, "La descripción del proyecto insertado no coincide con la esperada")
	// Puedes agregar más verificaciones según los campos que sean importantes verificar

	// Limpiar datos después de la prueba (eliminar el proyecto insertado)
	db, err := Connection_Migrates.Connect()
	if err != nil {
		t.Fatalf("Error conectando a la base de datos: %v", err)
	}
	err = Queries.DeleteProjectByID(db, insertedProject.IDProyecto)
	defer Connection_Migrates.Disconnect(db)

	if err != nil {
		t.Fatalf("Error limpiando datos después de la prueba: %v", err)
	}
}
