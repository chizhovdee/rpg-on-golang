package config
import (

	"github.com/jackc/pgx"
)

type Application struct {
	PgxConn *pgx.Conn
}

func NewApplication() *Application {
	return &Application{
		PgxConn: CreatePgxConn(),
	}
}