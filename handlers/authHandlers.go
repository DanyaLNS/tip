package handlers

import (
	"net/http"
	"tip/auth"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var creds auth.Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	// Здесь добавим простую проверку пароля
	if creds.Username != "user" || creds.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	token, err := auth.GenerateToken(creds.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

