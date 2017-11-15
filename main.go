package main

import (
  "log"
  "github.com/gin-gonic/gin"
  "github.com/hugomd/go-todo/db"
  TaskController "github.com/hugomd/go-todo/controllers"
)

func main() {
  log.Println("Starting server..")

  db.Init();

  r := gin.Default()

  v1 := r.Group("/api/v1")
  {
    tasks := v1.Group("/tasks")
    {
      tasks.GET("/", TaskController.GetTasks)
      tasks.POST("/", TaskController.CreateTask)
      tasks.PUT("/:id", TaskController.UpdateTask)
      tasks.DELETE("/:id", TaskController.DeleteTask)
    }
  }

  r.Run()
}
