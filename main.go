package main

import (
	"fmt"
	"go_todo/domain"
	"go_todo/infrastructure"
	"go_todo/infrastructure/middleware"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	Env_load()
	router := gin.Default()
	router.Use(middleware.RequestLogger())
	router.LoadHTMLGlob("templates/*/*")
	router.Static("/assets", "./assets")

	infrastructure.DBInit()

	router.GET("/", func(c *gin.Context) {
		todos := infrastructure.DBRead()
		c.HTML(200, "todos/index.html", gin.H{
			"todos": todos,
		})
	})

	router.GET("todos/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		todo := infrastructure.DBRead(id)
		c.HTML(200, "todos/show.html", gin.H{
			"todo": todo,
		})
	})

	router.POST("todos/new", func(c *gin.Context) {
		var todo domain.Todo
		err := c.BindJSON(&todo)
		fmt.Print(todo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		infrastructure.DBCreate(todo)
		c.Redirect(302, "/")
	})

	router.PUT("todos/:id", func(c *gin.Context) {
		text := c.PostForm("text")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		rawStatus := c.PostForm("status")
		statusID, err := strconv.Atoi(rawStatus)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		status := domain.Status(statusID)
		rawEstimate := c.PostForm("estimate")
		estimate, err := strconv.Atoi(rawEstimate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		rawTime := c.PostForm("time")
		var time int
		if rawTime != "" {
			time, err = strconv.Atoi(rawTime)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
		infrastructure.DBUpdate(id, text, status, estimate, time)

	})

	router.Run()
}
