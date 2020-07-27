package domain

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Text   string
	Status Status
	Time   int
}

type Status int

const (
	Task = iota
	ThisWeek
	Doing
	Review
	Done
	Close
)
