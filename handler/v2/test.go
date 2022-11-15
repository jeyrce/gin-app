package v2

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleTest(c *gin.Context) {
	c.JSON(
		http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "预留api版本",
			"data":    nil,
		},
	)
}
