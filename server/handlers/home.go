package handlers
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}