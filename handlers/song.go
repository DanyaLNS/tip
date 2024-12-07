package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"
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
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()  

	var songs []Song
	var total int64

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	query := db.DB.WithContext(ctx).Limit(limit).Offset(offset)

	fields := []string{"artist", "title", "album", "genre"}
	for _, field := range fields {
		value := c.Query(field)
		if value != "" {
			query = query.Where(field+" ILIKE ?", "%"+value+"%")
		}
	}

	if err := query.Find(&songs).Count(&total).Error; err != nil {
		handleError(c, http.StatusRequestTimeout, "Request timed out")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  songs,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}


func GetSongByID(c *gin.Context) {
	id := c.Param("id")
	var song Song
	if err := db.DB.First(&song, id).Error; err != nil {
		handleError(c, http.StatusNotFound, "song not found")
		return
	}
	c.JSON(http.StatusOK, song)
}

func CreateSong(c *gin.Context) {
	var newSong Song
	if err := c.BindJSON(&newSong); err != nil {
		handleError(c, http.StatusBadRequest, "invalid request")
		return
	}
	db.DB.Create(&newSong)
	c.JSON(http.StatusCreated, newSong)
}

func UpdateSong(c *gin.Context) {
	id := c.Param("id")
	var updatedSong Song
	if err := c.BindJSON(&updatedSong); err != nil {
		handleError(c, http.StatusBadRequest, "invalid request")
		return
	}
	if err := db.DB.Model(&Song{}).Where("id = ?", id).Updates(updatedSong).Error; err != nil {
		handleError(c, http.StatusNotFound, "song not found")
		return
	}
	c.JSON(http.StatusOK, updatedSong)
}

func DeleteSong(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&Song{}, id).Error; err != nil {
		handleError(c, http.StatusNotFound, "song not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "song deleted"})
}
