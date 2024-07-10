package routine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRoutines(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func DetailRoutine(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func AddRoutine(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func UpdateRoutine(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func ChangeRoutine(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func DeleteRoutine(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}
