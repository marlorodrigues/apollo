package authentication

import (
	"fmt"
	"net/http"
	"strings"

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

func ValidCookie(c *gin.Context) {
	cookie, err := c.Request.Cookie("_zd_a_a")

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not Authorized! Invalid Cookie."})
		return
	}

	value := strings.Replace(cookie.Value, "s:", "", 1)
	token := strings.Split(value, ".")

	c.Next()

	if validToken(token[0] + "." + token[1] + "." + token[2]) {
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Authorization!"})
}
