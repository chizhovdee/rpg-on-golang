package main
import (
	"github.com/gin-gonic/gin"
	"github.com/chizhovdee/rpg/server/handlers"
)

func setupRoutes(router *gin.Engine){
	router.Static("/assets", "./public/assets")
	router.Static("/images", "./public/images")
	router.Static("/locales", "./public/locales")
	//router.StaticFS("/assets", assetFS())

	router.LoadHTMLFiles("views/index.html")


	router.Use(SetCurrentCharacter())

	router.GET("/", handlers.HomeIndex)

	// API routers
	apiRouter := router.Group("/api/:version")

	apiRouter.GET("/characters/game_data.json", handlers.CharacterGameData)
	apiRouter.GET("/characters/status.json", handlers.CharacterStatus)

}