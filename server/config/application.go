package config
import "gopkg.in/gorp.v1"

type Application struct {
	DbMap *gorp.DbMap
}

func NewApplication() *Application {
	dbMap := InitDbMap()

	return &Application{DbMap: dbMap}
}