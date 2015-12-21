package tasks
import (
	"github.com/chuckpreslar/gofer"
	"html/template"
	"path/filepath"
	"os"
	"github.com/chizhovdee/rpg/server/game_data"
	"encoding/json"
)

var ExportGameData = gofer.Register(gofer.Task{
	Namespace:   "game_data",
	Label:       "export",
	Description: "Export game data to pure javascript",
	Action: func(arguments ...string) error {
		filename := "game_data.coffee"

		fpath := filepath.Join("../client/js/", filename)

		tplFuncs := template.FuncMap{
			"marshal": func(v interface {}) template.HTML {
				a, _ := json.Marshal(v)
				return template.HTML(a)
			},
		}
		t := template.Must(
			template.New("game_data.coffee.tpl").Funcs(tplFuncs).ParseFiles("tasks/game_data.coffee.tpl"),
		)

		f, e := os.Create(fpath)
		if e != nil {
			return e
		}
		defer f.Close()

		e = t.Execute(f, game_data.ExportToClient())

		if e != nil {
			return e
		}

	    return nil
	},
})