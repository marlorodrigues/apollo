package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pingcap/log"

	Authentication "apollo/src/handlers/authentication"
	Internals "apollo/src/handlers/internals"
)

func main() {
	log.Info("Starting Apollo...")

	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.HEAD("/status", Internals.APIStatus)

	v1.POST("/login", Authentication.Login)
	v1.POST("/logout", Authentication.Logout)

	// v1.POST("/tess", Routines.AddRoutine)

	router.Run("localhost:8080")
}
