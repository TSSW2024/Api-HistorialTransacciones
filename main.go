package main

import (
	"backend/api/config"
	"backend/api/handlers"
	_ "backend/api/models"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	// Conectar a la base de datos
	db := config.InitDatabase()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Crear el router
	r := gin.Default()
	// Configurar las rutas para CRUD de usuarios
	r.POST("/users", handlers.CreateUser(db))
	r.GET("/users/:id", handlers.GetUser(db))
	r.GET("/users", handlers.GetAllUsers(db))
	r.PUT("/users/:id", handlers.UpdateUser(db))
	r.DELETE("/users/:id", handlers.DeleteUser(db))

	//db.AutoMigrate(&models.LogEntry{})

	println("METODOS")
	println("Para PRobar usa: http://localhost:8084/")

	r.POST("/log", handlers.CreateLogs(db))
	//-----------------------------------------------------------
	r.GET("/log/todo", handlers.GetAllLogs(db))
	r.GET("/log/:UsuarioID", handlers.GetLog(db))
	//------------------------------------------------------------
	r.DELETE("/log/:UsuarioID", handlers.Deletelogs(db))
	r.POST("/wallet", handlers.CrearMonedero(db))
	r.GET("/wallet/:usuarioID", handlers.GetMonedero(db))
	// Iniciar el servidor
	r.Run(":8084")

}
