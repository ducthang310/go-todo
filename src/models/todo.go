package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Description string
	Complete    bool
}

type Checklist struct {
	Todo
	Description string
	Checked     bool
}
