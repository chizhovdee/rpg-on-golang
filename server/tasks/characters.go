package tasks
import (
	"github.com/chuckpreslar/gofer"
	"github.com/chizhovdee/rpg/server/models"
	"github.com/chizhovdee/rpg/server/config"
	"fmt"
)

var AddTestCharacter = gofer.Register(gofer.Task{
	Namespace:   "character",
	Label:       "add_test",
	Description: "Added one test character",
	Action: func(arguments ...string) error {
		dbMap := config.InitDbMap()

		ch := &models.Character{}

		err := dbMap.Insert(ch)

		if err != nil {
			return err
		}

		fmt.Println("Character added successfully!")

		return nil
	},
})