package main

import (
	"strconv"
	"time"

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

/*
Статусы:
todo
in_progress
done

Приоритеты:
low
medium
high

*/

var tasks = []Task{
	{
		ID:    1,
		Title: "Learn",
	},
}

func main() {
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(200, tasks)
	})

	r.POST("/tasks", func(ctx *gin.Context) {
		var task Task

		if err := ctx.BindJSON(&task); err != nil {
			ctx.JSON(400, gin.H{
				"error": "invalid json",
			})
			return
		}

		tasks = append(tasks, task)
		ctx.JSON(201, task)

	})

	r.DELETE("/tasks/:id", func(ctx *gin.Context) {
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

	r.GET("/tasks/:id", func(ctx *gin.Context) {
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

	r.PATCH("/tasks/:id", func(ctx *gin.Context) {
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
