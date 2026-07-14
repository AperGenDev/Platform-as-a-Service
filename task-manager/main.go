package main

import (
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	CreateAt    time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
}


var tasks = []Task{
	{
		ID:    1,
		Title: "Learn",
	},
}
var nextID = 2

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
    AllowOrigins: []string{"http://localhost:5173"},
    AllowMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
    AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	r.GET("/api/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.GET("/api/tasks", func(ctx *gin.Context) {
		ctx.JSON(200, tasks)
	})

	r.POST("/api/tasks", func(ctx *gin.Context) {
		var task Task

		if err := ctx.BindJSON(&task); err != nil {
			ctx.JSON(400, gin.H{
				"error": "invalid json",
			})
			return
		}
		if task.Title == "" {
			ctx.JSON(400, gin.H{
				"error": "title is required",
			})
			return
		}

		if task.Status != "" && !isValidStatus(task.Status) {
			ctx.JSON(400, gin.H{
				"error": "invalid status",
			})
			return
		}

		if task.Priority != "" && !isValidPriority(task.Priority) {
			ctx.JSON(400, gin.H{
				"error": "invalid priority",
			})
			return
		}

		task.ID = nextID
		nextID++

		task.CreateAt = time.Now()
		task.UpdateAt = time.Now()

		if task.Status == "" {
			task.Status = "todo"
		}

		if task.Priority == "" {
			task.Priority = "medium"
		}

		tasks = append(tasks, task)
		ctx.JSON(201, task)

	})

	r.DELETE("/api/tasks/:id", func(ctx *gin.Context) {
		taskID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "invalid id",
			})
			return
		}

		for i, task := range tasks {

			if task.ID == taskID {
				tasks = append(tasks[:i], tasks[i+1:]...)

				ctx.JSON(200, gin.H{
					"message": "task deleted",
				})

				return
			}
		}

		ctx.JSON(404, gin.H{
			"error": "task not found",
		})

	})

	r.GET("/api/tasks/:id", func(ctx *gin.Context) {
		taskID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "invalid id",
			})
			return
		}

		for i, task := range tasks {

			if taskID == task.ID {
				ctx.JSON(200, tasks[i])
				return
			}
		}

		ctx.JSON(404, gin.H{
			"error": "task not found",
		})
	})

	r.PATCH("/api/tasks/:id", func(ctx *gin.Context) {
		taskID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "invalid id",
			})
			return
		}

		var foundIndex = -1
		for i, task := range tasks {
			if taskID == task.ID {
				foundIndex = i
				break
			}
		}

		if foundIndex == -1 {
			ctx.JSON(404, gin.H{
				"error": "index not found",
			})
			return
		}

		var updates map[string]interface{}

		if err := ctx.BindJSON(&updates); err != nil {
			ctx.JSON(400, gin.H{
				"error": "invalid json",
			})
			return
		}

		if title, ok := updates["title"]; ok {
			tasks[foundIndex].Title = title.(string)
		}
		if description, ok := updates["description"]; ok {
			tasks[foundIndex].Description = description.(string)
		}
		if status, ok := updates["status"]; ok {
			tasks[foundIndex].Status = status.(string)
		}
		if priority, ok := updates["priority"]; ok {
			tasks[foundIndex].Priority = priority.(string)
		}

		tasks[foundIndex].UpdateAt = time.Now()

		ctx.JSON(200, tasks[foundIndex])
	})

	r.Run(":8080")
}
func isValidStatus(status string) bool {
	switch status {
	case "todo", "in_progress", "done":
		return true
	default:
		return false
	}
}

func isValidPriority(priority string) bool {
	switch priority {
	case "low", "medium", "high":
		return true
	default:
		return false
	}
}