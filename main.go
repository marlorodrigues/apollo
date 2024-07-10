package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pingcap/log"

	Authentication "apollo/src/handlers/authentication"
	Internals "apollo/src/handlers/internals"
	Routines "apollo/src/handlers/routines"
	Tags "apollo/src/handlers/tags"
	Tasks "apollo/src/handlers/tasks"
)

func main() {
	log.Info("Starting Apollo...")

	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.HEAD("/status", Internals.APIStatus)

	v1.POST("/login", Authentication.Login)
	v1.POST("/logout", Authentication.Logout)

	v1.GET("/tasks", Tasks.ListTasks)
	v1.GET("/task", Tasks.DetailTask)
	v1.POST("/task", Tasks.AddTask)
	v1.PUT("/task", Tasks.UpdateTask)
	v1.DELETE("/task", Tasks.DeleteTask)

	v1.GET("/tags", Tags.ListTags)
	v1.GET("/tag", Tags.DetailTag)
	v1.POST("/tag", Tags.AddTag)
	v1.PUT("/tag", Tags.UpdateTag)
	v1.DELETE("/tag", Tags.DeleteTag)

	v1.GET("/routines", Routines.ListRoutines)
	v1.GET("/routine", Routines.DetailRoutine)
	v1.POST("/routine", Routines.AddRoutine)
	v1.PUT("/routine", Routines.UpdateRoutine)
	v1.DELETE("/routine", Routines.DeleteRoutine)

	// v1.POST("/tess", Routines.AddRoutine)

	router.Run("localhost:8080")
}
