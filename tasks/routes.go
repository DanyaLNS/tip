package tasks

import (
	"net/http"
	"tip/db"
	"tip/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/tasks", func(c *gin.Context) {
		var newSong handlers.Song
		if err := c.BindJSON(&newSong); err != nil {
			handlers.HandleError(c, http.StatusBadRequest, "invalid song data")
		}

		if err := db.DB.Create(&newSong).Error; err != nil {
			handlers.HandleError(c, http.StatusInternalServerError, "failed to save song")
			return
		}

		taskID := CreateTask(newSong.ID)
		go RunTask(taskID)
		c.JSON(201, gin.H{"task_id": taskID})
	})

	r.GET("/tasks/:id", func(c *gin.Context) {
		taskID := c.Param("id")
		task := GetTask(taskID)
		if task == nil {
			handlers.HandleError(c, http.StatusNotFound, "task not found")
		}
		c.JSON(200, task)
	})
}
