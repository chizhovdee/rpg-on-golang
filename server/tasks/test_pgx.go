package tasks
import (
	"github.com/chuckpreslar/gofer"
	"github.com/jackc/pgx"
	"fmt"
	log "gopkg.in/inconshreveable/log15.v2"
)

var conn *pgx.Conn

type Character struct {
	fields map[string]interface{}
}


var TestPgx = gofer.Register(gofer.Task{
	Label:       "pgx",
	Description: "Test PGX",
	Action: func(arguments ...string) error {

		var err error
		conn, err = pgx.Connect(extractConfig())

		if err != nil {
			return err
		}

		ch := &Character{fields: map[string]interface{}{}}

		rows, err := conn.Query("select * from characters where id=$1 limit 1", 1)

		for rows.Next() {
			values, _ := rows.Values()

			for index, field := range rows.FieldDescriptions() {
				ch.fields[field.Name] = values[index]
				//fmt.Println(field.Name, values[index])
			}
		}

		fmt.Println("ROW", ch.fields)

		return nil
	},
})

func extractConfig() pgx.ConnConfig {
	var config pgx.ConnConfig

	config.Host = "localhost"

	config.User = "deewild"

	config.Password = "satch"

	config.Database = "rpg_game_development"

	config.LogLevel = pgx.LogLevelInfo

	config.Logger = log.New("DBdriver", "Postgresql")

	return config
}