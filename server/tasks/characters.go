package tasks
import (
	"github.com/chuckpreslar/gofer"
	"github.com/chizhovdee/rpg/server/config"
	"fmt"
)

var AddTestCharacter = gofer.Register(gofer.Task{
	Namespace:   "character",
	Label:       "add",
	Description: "Added one test character",
	Action: func(arguments ...string) error {
		var err error
		conn := config.CreatePgxConn()

		_, err = conn.Exec(`
		insert into characters(energy, ep, hp, health, basic_money, vip_money)
		                values($1, $2, $3, $4, $5, $6)
		`, 10, 10, 100, 100, 100, 1)

		if err != nil {
			return err
		}

		fmt.Println("Character added successfully!")

		return nil
	},
})