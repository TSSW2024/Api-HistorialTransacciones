package main

import (
	"backend/api/handlers"
	_ "backend/api/models"
	"backend/api/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar a la base de datos
	db, err := utils.OpenGormDB()
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
	r.GET("/log/todo", handlers.GetAllLogs(db))   //ahora si funca
	r.GET("/log/:UsuarioID", handlers.GetLog(db)) //no funca
	//------------------------------------------------------------
	r.DELETE("/log/:UsuarioID", handlers.Deletelogs(db))

	// Iniciar el servidor
	r.Run(":8084")

}
