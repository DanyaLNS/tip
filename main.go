package main

import (
	"tip/db"
	"tip/handlers"
	"tip/tasks"

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

	handlers.SetupRoutes(r)
	tasks.SetupRoutes(r)

	r.Run()
}
