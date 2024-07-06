package handlers

import (
	"backend/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Deletelogs elimina un usuario de la base de datos
func Deletelogs(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("UsuarioDentro")
		var log models.LogEntry
		result := db.First(&log, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User no se encuentra en la base de datos"})
			return
		}

		db.Delete(&log)
		c.JSON(http.StatusNoContent, nil)
	}
}
