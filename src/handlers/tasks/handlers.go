package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func DetailTask(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func AddTask(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func UpdateTask(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func ChangeTask(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func DeleteTask(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}
