package v1

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	
	"github.com/jeyrce/gin-app/pkg/conf"
)

func handleVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "system version",
		"data":    conf.MetaVersionMap,
	})
}
