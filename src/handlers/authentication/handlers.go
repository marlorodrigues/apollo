package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func Logout(c *gin.Context) {
	var loginData sLogin

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	is_valid := validLogin(&loginData)

	if is_valid != "" {

	}

	c.IndentedJSON(http.StatusOK, gin.H{})
}
