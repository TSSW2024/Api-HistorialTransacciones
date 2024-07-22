package handlers

import (
	"backend/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetLog(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		UsuarioID := c.Param("UsuarioID")

		// Buscar todos los logs basados en id_session
		var logs []models.LogEntry
		if err := db.Preload("Usuario").Where("id_session = ?", UsuarioID).Find(&logs).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "No logs found for the given id_session"})
			return
		}

		c.JSON(http.StatusOK, logs)
	}
}

func GetAllLogs(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var logs []models.LogEntry
		if err := db.Find(&logs).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, logs)
	}
}
