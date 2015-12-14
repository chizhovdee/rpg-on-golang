package tasks
import (
	"github.com/chuckpreslar/gofer"
	"html/template"
	"path/filepath"
	"os"
)

var ExportGameData = gofer.Register(gofer.Task{
	Namespace:   "game_data",
	Label:       "export",
	Description: "Export game data to pure javascript",
	Action: func(arguments ...string) error {
		filename := "game_data.js"

		fpath := filepath.Join(".", filename)
		t := exportGameDataTemplate

		f, e := os.Create(fpath)
		if e != nil {
			return e
		}
		defer f.Close()

		e = t.Execute(f, map[string]interface{}{"missions": 1})

		if e != nil {
			return e
		}


	    return nil
	},
})

var exportGameDataTemplate = template.Must(template.New("game-data-template").Parse(
`
console.log({{.missions}});
`))