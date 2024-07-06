package handlers

import (
	"backend/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetLog(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		UsuarioID := c.Param("UsuarioID")

		// Convertir UsuarioID a uint
		id, err := strconv.ParseUint(UsuarioID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UsuarioID"})
			return
		}

		var log models.LogEntry
		if err := db.Preload("Usuario").First(&log, "usuario_id = ?", uint(id)).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Log entry not found"})
			return
		}

		c.JSON(http.StatusOK, log)
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
