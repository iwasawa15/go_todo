package domain

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Text     string `json:"text"`
	Status   Status `json:"status"`
	Estimate int    `json:"estimate"`
	Time     int    `json:"time"`
}

type Status int

const (
	Task Status = iota
	ThisWeek
	Doing
	Review
	Done
	Close
)
