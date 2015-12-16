package tasks
import (
	"github.com/chuckpreslar/gofer"
	"html/template"
	"path/filepath"
	"os"
	"github.com/chizhovdee/rpg/server/game_data"
	"encoding/json"
)

func gameData() map[string]interface{} {
	game_data.Define()

	return map[string]interface{}{
		"missions": game_data.Missions.AsJson(),
	}
}

var ExportGameData = gofer.Register(gofer.Task{
	Namespace:   "game_data",
	Label:       "export",
	Description: "Export game data to pure javascript",
	Action: func(arguments ...string) error {
		filename := "game_data.coffee"

		fpath := filepath.Join(".", filename)

		tplFuncs := template.FuncMap{
			"marshal": func(v interface {}) template.HTML {
				a, _ := json.Marshal(v)
				return template.HTML(a)
			},
		}
		t := template.Must(template.New("game_data.coffee").Funcs(tplFuncs).ParseFiles("tasks/game_data.coffee"))

		f, e := os.Create(fpath)
		if e != nil {
			return e
		}
		defer f.Close()

		e = t.Execute(f, gameData())

		if e != nil {
			return e
		}


	    return nil
	},
})