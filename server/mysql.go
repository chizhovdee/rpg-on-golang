package main
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"log"
)

var DbMap *gorp.DbMap

func CreateDbMap() {
	db, err := sql.Open("mysql", "root:@/rpg_game_development?parseTime=true")

	if err != nil {
		log.Fatalln(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	DbMap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
}

