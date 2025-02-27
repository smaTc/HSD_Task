package main

import (
	"strconv"
	"strings"

	"github.com/smaTc/HSD_Task/hsdtask"

	"github.com/gin-gonic/gin"
)

var logger = hsdtask.Logger
var port = "8000"
var dbPath = "hsd.db"

func main() {
	hsdtask.InitDB(dbPath)

	router := gin.Default()
	router.GET("/tasks/*id", func(c *gin.Context) {
		id := c.Param("id")
		id = strings.Replace(id, "/", "", -1)

		if id == "" {
			c.Writer.Write(hsdtask.GetTasks())
		} else {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				msg := "invalid parameter for task id"
				logger.Println(msg)
				c.Writer.Write(hsdtask.GenerateJsonResponse("error", msg))
			} else {
				c.Writer.Write(hsdtask.GetTasks(idInt))
			}
		}
	})

	router.POST("/tasks", func(c *gin.Context) {
		var task hsdtask.Task
		err := c.ShouldBindJSON(&task)
		if err != nil {
			c.Writer.Write(hsdtask.GenerateJsonResponse("error", "creating task failed do to invalid payload/JSON"))
		} else {
			c.Writer.Write(hsdtask.CreateTask(&task))
		}
	})

	router.PUT("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")
		id = strings.Replace(id, "/", "", -1)

		var task hsdtask.Task
		err := c.ShouldBindJSON(&task)

		if id == "" {
			c.Writer.Write(hsdtask.GenerateJsonResponse("error", "no id for task update provided"))
		} else if err != nil {
			c.Writer.Write(hsdtask.GenerateJsonResponse("error", "creating task failed do to invalid payload/JSON"))
		} else {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				msg := "invalid parameter for task id"
				logger.Println(msg)
				c.Writer.Write(hsdtask.GenerateJsonResponse("error", msg))
			} else {
				c.Writer.Write(hsdtask.UpdateTask(idInt, &task))
			}
		}
	})

	router.DELETE("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")
		id = strings.Replace(id, "/", "", -1)

		if id == "" {
			c.Writer.Write(hsdtask.GenerateJsonResponse("error", "no id for task deletion provided"))
		} else {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				msg := "invalid parameter for task id"
				logger.Println(msg)
				c.Writer.Write(hsdtask.GenerateJsonResponse("error", msg))
			} else {
				c.Writer.Write(hsdtask.DeleteTask(idInt))
			}
		}
	})

	logger.Println("running on port", port)
	router.Run(":" + port)
}
