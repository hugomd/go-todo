package task

import (
  "net/http"
  "github.com/gin-gonic/gin"
  // "github.com/gin-gonic/gin/binding"
  "github.com/hugomd/go-todo/models"
  "github.com/hugomd/go-todo/db"
)

func GetTasks(c *gin.Context) {

  var tasks []models.Task
  db := db.GetDB()
  db.Find(&tasks)
  c.JSON(200, tasks)
}

func CreateTask(c *gin.Context) {
  var task models.Task
  var db = db.GetDB()

  if err := c.BindJSON(&task); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }
  db.Create(&task)
  c.JSON(http.StatusOK, &task)
}

func UpdateTask(c *gin.Context) {
  id := c.Param("id")
  var task models.Task

  db := db.GetDB()
  if err := db.Where("id = ?", id).First(&task).Error; err != nil {
    c.AbortWithStatus(http.StatusNotFound)
    return
  }
  c.BindJSON(&task)
  db.Save(&task)
  c.JSON(http.StatusOK, &task)
}

func DeleteTask(c *gin.Context) {
  id := c.Param("id")
  var task models.Task
  db := db.GetDB()

  if err := db.Where("id = ?", id).First(&task).Error; err != nil {
    c.AbortWithStatus(http.StatusNotFound)
    return
  }

	db.Delete(&task)
}
