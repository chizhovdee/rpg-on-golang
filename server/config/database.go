package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"gopkg.in/gorp.v1"
	"os"
	"github.com/chizhovdee/rpg/server/models"
	"log"
)

// Структура конфигурации базы данных для чтения из yaml файла
type dbConfig struct {
	Driver string `yaml:"driver"`
	Dialect string `yaml:"dialect"`
	Database string `yaml:"database"`
	Host string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	MigrationsDir string `yaml:"migrations_dir"`
}

// Устанавливает соединение с базой
func OpenDb() (*sql.DB, error){
	config, err := readConfig()

	if err != nil {
		return nil, err
	}

	conf := config[os.Getenv("ENV")]

	db, err := sql.Open(
		"mysql",
		conf.Username + ":" + conf.Password + "@/" + conf.Database  + "?parseTime=true",
	)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// Инициализируем ОРМ
func InitDbMap() *gorp.DbMap {
	db, err := OpenDb()

	if err != nil {
		panic(err.Error())
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	setTables(dbMap)

	dbMap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	return dbMap
}

// Читает yml файл с конфигурацией базы данных
func readConfig() (map[string]*dbConfig, error) {
	file, err := ioutil.ReadFile("config/database.yml")
	if err != nil {
		return nil, err
	}

	config := make(map[string]*dbConfig)
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func setTables(dbMap *gorp.DbMap) {
	dbMap.AddTableWithName(models.Character{}, "characters").SetKeys(true, "Id")
}