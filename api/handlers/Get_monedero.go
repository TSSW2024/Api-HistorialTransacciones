package handlers

import (
	"backend/api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetMonedero(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarioID := c.Param("usuarioID")
		var monedero models.Monedero

		// Debugging log
		log.Printf("Buscando monedero para usuarioID: %s", usuarioID)

		if err := db.Preload("Monedas").Where("usuario_id = ?", usuarioID).First(&monedero).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "monedero no encontrado"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error al obtener el monedero"})
			}
			return
		}

		c.JSON(http.StatusOK, monedero)
	}
}
