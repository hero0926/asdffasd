package basic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Main Page
func Index(c *gin.Context) {

	c.HTML(http.StatusOK, "basic/main", gin.H{
		"title":    "Main",
		"articles": "articles",
	})
}
