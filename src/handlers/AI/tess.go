package ai

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListChats(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func DetailChat(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func AddChat(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func UpdateChat(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func ChangeChat(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func DeleteChat(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}
