package handlers
import "github.com/chizhovdee/rpg/server/config"

var App *config.Application

func SetupApplication(app *config.Application){
	App = app
}