package tasks

import (
    "encoding/json"
    "os"
    "sync"
    "time"
    "tip/db"
    "tip/handlers"
)

type Task struct {
    ID       string `json:"id"`
    Status   string `json:"status"`   
    Filename string `json:"filename"`
    SongID   int    `json:"song_id"`
}

var tasks = make(map[string]*Task)
var mutex = &sync.Mutex{}

func CreateTask(songID int) string {
    taskID := generateID()
    task := &Task{
        ID:     taskID,
        Status: "pending",
        SongID: songID,
    }
    mutex.Lock()
    tasks[taskID] = task
    mutex.Unlock()
    return taskID
}

func GetTask(taskID string) *Task {
    mutex.Lock()
    defer mutex.Unlock()
    return tasks[taskID]
}

func RunTask(taskID string) {
    task := GetTask(taskID)
    if task == nil {
        return
    }

    task.Status = "in_progress"
    filename := "export_song_" + task.ID + ".json"
    task.Filename = filename

    var song handlers.Song
    if err := db.DB.First(&song, task.SongID).Error; err != nil {
        task.Status = "error"
        return
    }

    time.Sleep(5 * time.Second) // Симуляция выполнения задачи

    file, err := os.Create(filename)
    if err != nil {
        task.Status = "error"
        return
    }
    defer file.Close()

    data := map[string]interface{}{
        "title":  song.Title,
        "artist": song.Artist,
        "album":  song.Album,
        "genre":  song.Genre,
        "duration": song.Duration,
    }
    json.NewEncoder(file).Encode(data)

    task.Status = "done"
}

func generateID() string {
    return time.Now().Format("20060102150405")
}
