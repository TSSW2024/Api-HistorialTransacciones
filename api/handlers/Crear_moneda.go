package handlers

import (
	"backend/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MonedaRequest struct {
	UserID string `json:"userId"`
	Moneda struct {
		ID       uint   `json:"id"`
		Nombre   string `json:"nombre"`
		Cantidad int    `json:"cantidad"`
	} `json:"moneda"`
}

// GetUser devuelve un gin.HandlerFunc que maneja la creación de una moneda
func CrearMoneda(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req MonedaRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Buscar el monedero del usuario
		var monedero models.Monedero
		if err := db.Where("usuario_id = ?", req.UserID).First(&monedero).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Monedero not found for user"})
			return
		}

		// Crear la nueva moneda
		nuevaMoneda := models.Moneda{
			ID:         req.Moneda.ID,
			Nombre:     req.Moneda.Nombre,
			Cantidad:   req.Moneda.Cantidad,
			MonederoID: monedero.ID,
		}

		// Agregar la moneda a la base de datos
		if err := db.Create(&nuevaMoneda).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create moneda"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Moneda created successfully"})
	}
}
