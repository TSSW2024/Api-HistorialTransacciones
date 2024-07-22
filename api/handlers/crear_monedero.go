package handlers

import (
	"backend/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CrearMonedero maneja la creación de un monedero vacío si no existe
func CrearMonedero(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			UsuarioID string `json:"usuarioID" binding:"required"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var existingMonedero models.Monedero
		if err := db.Where("usuario_id = ?", input.UsuarioID).First(&existingMonedero).Error; err != nil && err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking for existing monedero"})
			return
		}

		if existingMonedero.ID != 0 {
			// Monedero existe, devolvemos el existente
			c.JSON(http.StatusOK, existingMonedero)
		} else {
			// Monedero no existe, creamos uno nuevo vacío
			newMonedero := models.Monedero{
				UsuarioID: input.UsuarioID,
				Monedas:   []models.Moneda{},
			}
			if err := db.Create(&newMonedero).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating monedero" + err.Error()})
				return
			}
			c.JSON(http.StatusCreated, newMonedero)
		}
	}
}
