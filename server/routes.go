package main
import "github.com/gin-gonic/gin"

func setupRoutes(router *gin.Engine){
	router.Static("/assets", "./assets")
	//router.StaticFS("/assets", assetFS())

	router.LoadHTMLFiles("views/index.html")

	router.GET("/", index)
}
