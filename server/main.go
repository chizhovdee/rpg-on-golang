package main
import (
	"github.com/gin-gonic/gin"
	"os"
	"github.com/chizhovdee/rpg/server/models"
	"github.com/chizhovdee/rpg/server/handlers"
	"github.com/chizhovdee/rpg/server/config"
	"github.com/chizhovdee/rpg/server/game_data"
)

func init() {
	if os.Getenv("ENV") == "" {
		os.Setenv("ENV", "development")
	}

	if os.Getenv("PORT") == "" && os.Getenv("ENV") == "development" {
		os.Setenv("PORT", "3000")
	}
}

func main(){
	app := config.NewApplication()

	handlers.SetupApplication(app)

	models.SetupApplication(app)

	game_data.Define()

	router := gin.Default()

	setupRoutes(router)

	router.Run(":" + os.Getenv("PORT"))
}
