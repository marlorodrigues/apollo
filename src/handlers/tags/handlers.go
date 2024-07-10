package tags

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTags(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func DetailTag(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func AddTag(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func UpdateTag(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func ChangeTag(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func DeleteTag(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}
