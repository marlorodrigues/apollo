package internals

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}
