package db

import (
    "fmt"
    "log"
  _ "github.com/lib/pq"
    "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/hugomd/go-todo/models"
)

var db *gorm.DB
var err error

// Init creates a connection to mysql database and 
// migrates any new models
func Init() {
  dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
    "hugo",
    "",
    "localhost",
    "5432",
    "tasks",
  )


  db, err = gorm.Open("postgres", dbinfo)
  if err != nil {
    log.Println("Failed to connect to database")
    panic(err)
  }
  log.Println("Database connected")

  if !db.HasTable(&models.Task{}) {
    err := db.CreateTable(&models.Task{})
    if err != nil {
      log.Println("Table already exists")
    }
  }

  db.AutoMigrate(&models.Task{})
}

//GetDB ...
func GetDB() *gorm.DB {
  return db
}

func CloseDB() {
  db.Close()
}
