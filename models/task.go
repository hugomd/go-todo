package models

import (
  "github.com/satori/go.uuid"
  "github.com/jinzhu/gorm"
)

// Task "Object
type Task struct {
  ID  uuid.UUID   `json:"id"`
  Title string  `json:"title"`
  CreatedAt string  `json:"created_at"`
  Completed bool  `json:"completed"`
}

func (task *Task) BeforeCreate(scope *gorm.Scope) error {
    return scope.SetColumn("ID", uuid.NewV4().String())
}
