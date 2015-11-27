package main
import (
	"github.com/gin-gonic/gin"
	"os"
	"github.com/chizhovdee/rpg/server/config"
	"github.com/chizhovdee/rpg/server/handlers"
	"fmt"
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

	router := gin.Default()

	setupRoutes(router)

	router.Run(":" + os.Getenv("PORT"))
}
