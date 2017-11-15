package models

import (
  "time"
  "github.com/satori/go.uuid"
  "github.com/jinzhu/gorm"
)

// Task "Object
type Task struct {
  ID  uuid.UUID   `json:"id"`
  Title string  `json:"title" binding:"required"`
  CreatedAt time.Time  `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  Completed bool  `json:"completed"`
}

func (task *Task) BeforeCreate(scope *gorm.Scope) error {
  scope.SetColumn("CreatedAt", time.Now())
  scope.SetColumn("ID", uuid.NewV4().String())
  return nil
}

func (task *Task) BeforeUpdate(scope *gorm.Scope) error {
  scope.SetColumn("UpdatedAt", time.Now())
  return nil
}
