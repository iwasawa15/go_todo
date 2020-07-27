package infrastructure

import (
	"fmt"
	"go_todo/domain"

	"github.com/jinzhu/gorm"
)

func DbInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/go_todo")
	if err != nil {
		fmt.Errorf("could not open database")
	}
	db.AutoMigrate(&domain.Todo{})
	return db
}

func DbCreate(text string, status domain.Status, time int) {
	db, err := gorm.Open("mysql", "root:@/go_todo")
	if err != nil {
		fmt.Errorf("could not open database")
	}
	db.Create(&domain.Todo{Text: text, Status: status, Time: time})
}

func DbRead(id int) []domain.Todo {
	db, err := gorm.Open("mysql", "root:@/go_todo")
	if err != nil {
		fmt.Errorf("could not open database")
	}
	var todos []domain.Todo
	db.Find(&todos)
	return todos
}

func DbUpdate(id int, text string, status domain.Status, time int) domain.Todo {
	db, err := gorm.Open("mysql", "root:@/go_todo")
	if err != nil {
		fmt.Errorf("could not open database")
	}
	var todo domain.Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	todo.Time = time
	db.Save(&todo)
	return todo
}

func DbDelete(id int) {
	db, err := gorm.Open("mysql", "root:@/go_todo")
	if err != nil {
		fmt.Errorf("could not open database")
	}
	var todo domain.Todo
	db.First(&todo, id)
	db.Delete(&todo)
}
