package handlers

import (
	"net/http"
	"tip/db"

	"github.com/gin-gonic/gin"
)

type Song struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Genre    string `json:"genre"`
	Duration string `json:"duration"`
}

func GetSongs(c *gin.Context) {
	var songs []Song
	db.DB.Find(&songs)
	c.JSON(http.StatusOK, songs)
}

func GetSongByID(c *gin.Context) {
	id := c.Param("id")
	var song Song
	if err := db.DB.First(&song, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "song not found"})
		return
	}
	c.JSON(http.StatusOK, song)
}

func CreateSong(c *gin.Context) {
	var newSong Song
	if err := c.BindJSON(&newSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	db.DB.Create(&newSong)
	c.JSON(http.StatusCreated, newSong)
}

func UpdateSong(c *gin.Context) {
	id := c.Param("id")
	var updatedSong Song
	if err := c.BindJSON(&updatedSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	if err := db.DB.Model(&Song{}).Where("id = ?", id).Updates(updatedSong).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "song not found"})
		return
	}
	c.JSON(http.StatusOK, updatedSong)
}

func DeleteSong(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&Song{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "song not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "song deleted"})
}
