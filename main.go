package main

import (
	"tip/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/songs", handlers.GetSongs)
	r.GET("/songs/:id", handlers.GetSongByID)
	r.GET("/albums", handlers.GetAlbums)
	r.GET("/albums/:id", handlers.GetAlbumByID)
	r.GET("/artists", handlers.GetArtists)
	r.GET("/artists/:id", handlers.GetArtistByID)

	r.Run()
}
