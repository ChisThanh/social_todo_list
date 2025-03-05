package model

import (
	"fmt"
	"social_todo_list/common"
)

type ItemStatus string

var (
	ErrTitleEmpty   = fmt.Errorf("title is empty")
	ErrItemNotFound = fmt.Errorf("item not found")
	ErrItemDeleted  = fmt.Errorf("item is deleted")
)

const (
	StatusDoing  ItemStatus = "Doing"
	StatusDone   ItemStatus = "Done"
	StatusDelete ItemStatus = "Deleted"
)

type TodoItems struct {
	common.SQLModel
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      *ItemStatus `json:"status"`
	Image       string      `json:"image,omitempty"`
}

type TodoItemCreation struct {
	Id          int         `json:"id" gorm:"column:id"`
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status,omitempty" gorm:"column:status"`
}

type TodoItemUpdate struct {
	Title       *string     `json:"title" gorm:"column:title"`
	Description *string     `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status,omitempty" gorm:"column:status"`
}

func (TodoItems) TableName() string        { return "todo_items" }
func (TodoItemCreation) TableName() string { return TodoItems{}.TableName() }
func (TodoItemUpdate) TableName() string   { return TodoItems{}.TableName() }
