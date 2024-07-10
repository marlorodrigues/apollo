package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func Logout(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}
