package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	gin.Logger()

	router.Static("/assets", "./assets")

	router.LoadHTMLFiles("views/index.html")

	router.GET("/", index)

	router.Run(":3000")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
