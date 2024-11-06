package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SongResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Genre    string `json:"genre"`
	Duration string `json:"duration"`
}

type AlbumResponse struct {
	ID     int            `json:"id"`
	Title  string         `json:"title"`
	Artist string         `json:"artist"`
	Year   int            `json:"year"`
	Genre  string         `json:"genre"`
	Songs  []SongResponse `json:"songs,omitempty"`
}

type ArtistResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Genre   string `json:"genre"`
	Country string `json:"country"`
}

var (
	Songs = []SongResponse{
		{ID: 1, Title: "Karma Police", Artist: "Radiohead", Album: "OK Computer", Genre: "Alternative Rock", Duration: "4:21"},
		{ID: 2, Title: "No Surprises", Artist: "Radiohead", Album: "OK Computer", Genre: "Alternative Rock", Duration: "3:48"},
	}

	Albums = []AlbumResponse{
		{ID: 1, Title: "OK Computer", Artist: "Radiohead", Year: 1997, Genre: "Alternative Rock", Songs: Songs},
		{ID: 2, Title: "In Rainbows", Artist: "Radiohead", Year: 2007, Genre: "Alternative Rock"},
	}

	Artists = []ArtistResponse{
		{ID: 1, Name: "Radiohead", Genre: "Alternative Rock", Country: "UK"},
	}
)

func GetSongs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": Songs})
}

func GetSongByID(c *gin.Context) {
	songID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	for _, song := range Songs {
		if songID == song.ID {
			c.JSON(http.StatusOK, gin.H{"result": song})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
}

func GetArtists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": Artists})
}

func GetArtistByID(c *gin.Context) {
	artistID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	for _, artist := range Artists {
		if artistID == artist.ID {
			c.JSON(http.StatusOK, gin.H{"result": artist})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Artist not found"})
}

func GetAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": Albums})
}

func GetAlbumByID(c *gin.Context) {
	albumID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	for _, album := range Albums {
		if albumID == album.ID {
			c.JSON(http.StatusOK, gin.H{"result": album})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
}
