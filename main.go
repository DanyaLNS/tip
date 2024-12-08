package main

import (
	"tip/auth"
	"tip/db"
	"tip/handlers"

	"github.com/gin-gonic/gin"
)
func setupDB() {
	db.InitDB()
	db.DB.AutoMigrate(&handlers.Song{})
}

func main() {
	setupDB()

	r := gin.Default()
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(auth.AuthMiddleware())
	{
		protected.POST("/songs", handlers.CreateSong)
		protected.PUT("/songs/:id", handlers.UpdateSong)
		protected.DELETE("/songs/:id", handlers.DeleteSong)
	}

	r.GET("/songs", handlers.GetSongs)
	r.GET("/songs/duration", handlers.GetSongsByDuration)
	r.GET("/songs/:id", handlers.GetSongByID)
	r.GET("/songs/by-artist", handlers.CountSongsByArtist)
	r.PUT("/songs/:id/label", handlers.UpdateSongsLabel)


	r.Run()
}
