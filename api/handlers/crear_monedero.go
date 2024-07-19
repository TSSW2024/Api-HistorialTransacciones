package handlers

import (
	"backend/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CrearMonedero maneja la creación o actualización del monedero
func CrearMonedero(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Monedero
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
			// Monedero existe, actualizamos
			db.Model(&existingMonedero).Updates(input)
			c.JSON(http.StatusOK, existingMonedero)
		} else {
			// Monedero no existe, creamos uno nuevo
			if err := db.Create(&input).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating monedero"})
				return
			}
			c.JSON(http.StatusCreated, input)
		}
	}
}
