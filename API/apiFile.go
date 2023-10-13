package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Task struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed" `
	Date      string `json:"date"`
}

var tasks = []Task{
	{ID: "1", Item: "Clean a room", Completed: false, Date: "20.02.2020"},
	{ID: "2", Item: "Go shopping", Completed: false, Date: "22.02.2020"},
	{ID: "3", Item: "Cook dinner", Completed: false, Date: "22.02.2020"},
}

func getTasks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, tasks)
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	err := router.Run("localhost:9090")
	if err != nil {
		return
	}
}
