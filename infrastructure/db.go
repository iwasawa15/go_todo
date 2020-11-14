package infrastructure

import (
	"database/sql"
	"fmt"
	"go_todo/domain"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDB() *sql.DB {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("MYSQL_HOST")
	sqlDB, err := sql.Open("mysql", user+":"+password+"@tcp("+host+")/"+database+"?parseTime=true&loc=Local")
	if err != nil {
		fmt.Errorf("could not open database")
	}
	return sqlDB
}

func DBInit() *gorm.DB {
	sqlDB := connectDB()
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Errorf("could not open database")
	}
	db.AutoMigrate(&domain.Todo{})
	return db
}

func DBCreate(todo domain.Todo) {
	sqlDB := connectDB()
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Errorf("could not open database")
	}
	db.Create(&todo)
}

func DBRead(id ...int) []domain.Todo {
	sqlDB := connectDB()
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Errorf("could not open database")
	}
	var todos []domain.Todo
	db.Find(&todos)
	return todos
}

func DBUpdate(id int, text string, status domain.Status, estimate int, time int) domain.Todo {
	sqlDB := connectDB()
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Errorf("could not open database")
	}
	var todo domain.Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	todo.Estimate = estimate
	todo.Time = time
	db.Save(&todo)
	return todo
}

func DBDelete(id int) {
	sqlDB := connectDB()
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Errorf("could not open database")
	}
	var todo domain.Todo
	db.First(&todo, id)
	db.Delete(&todo)
}
