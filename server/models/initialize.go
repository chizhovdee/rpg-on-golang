package models
import "github.com/chizhovdee/rpg/server/config"

var App *config.Application

func SetupApplication(app *config.Application){
	App = app
}

func SetupTables() {
	App.DbMap.AddTableWithName(Character{}, "characters").SetKeys(true, "Id")
}