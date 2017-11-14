package task

import (
  "log"
  "github.com/gin-gonic/gin"
  "github.com/hugomd/go-todo/models"
  "github.com/hugomd/go-todo/db"
)

func GetTasks(c *gin.Context) {

  var tasks []models.Task
  db := db.GetDB()
  db.Find(&tasks)
  c.JSON(200, tasks)
  return
}

func CreateTask(c *gin.Context) {
  var task models.Task
  var db = db.GetDB()

  err := c.BindJSON(&task)
  if err != nil {
    log.Println(err)
    c.Error(err).SetType(gin.ErrorTypeBind)
    return
  }
  db.Create(&task)
  c.JSON(200, task)
  return
}
