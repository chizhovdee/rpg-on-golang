package main
import (
	"github.com/gin-gonic/gin"
	"os"
	"github.com/chizhovdee/rpg/server/config"
	"github.com/chizhovdee/rpg/server/handlers"
)

func init() {
	if os.Getenv("ENV") == "" {
		os.Setenv("ENV", "development")
	}
}

func main(){
	app := config.NewApplication()

	handlers.SetupApplication(app)

	router := gin.Default()

	setupRoutes(router)

	router.Run(":3000")
}
