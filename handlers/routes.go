package handlers

import (
	"tip/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	protected := r.Group("/")
	protected.Use(auth.AuthMiddleware())
	{
		protected.POST("/songs", CreateSong)
		protected.PUT("/songs/:id", UpdateSong)
		r.PUT("/songs/:id/label", UpdateSongsLabel)
		protected.DELETE("/songs/:id", DeleteSong)
	}

	r.GET("/songs", GetSongs)
	r.GET("/songs/:id", GetSongByID)
	r.GET("/songs/duration", GetSongsByDuration)
	r.GET("/songs/by-artist", CountSongsByArtist)
}