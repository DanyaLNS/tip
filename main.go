package main

import (
	"tip/auth"
	"tip/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("/songs", handlers.GetSongs)
		protected.GET("/songs/:id", handlers.GetSongByID)
		protected.GET("/albums", handlers.GetAlbums)
		protected.GET("/albums/:id", handlers.GetAlbumByID)
		protected.GET("/artists", handlers.GetArtists)
		protected.GET("/artists/:id", handlers.GetArtistByID)
	}

	r.Run()
}
