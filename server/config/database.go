package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/jackc/pgx"
	log "gopkg.in/inconshreveable/log15.v2"
	"os"
)

// Структура конфигурации базы данных для чтения из yml файла
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
func CreatePgxConn() *pgx.Conn{
	var err error
	var pgxConfig pgx.ConnConfig
	var conn *pgx.Conn

	configYml, err := readConfig()

	if err != nil {
		panic(err.Error())
	}

	config := configYml[os.Getenv("ENV")]

	pgxConfig.Host = config.Host
	pgxConfig.User = config.Username
	pgxConfig.Password = config.Password
	pgxConfig.Database = config.Database
	pgxConfig.LogLevel = pgx.LogLevelInfo
	pgxConfig.Logger = log.New("DBdriver", "Postgresql")

	conn, err = pgx.Connect(pgxConfig)

	if err != nil {
		panic(err.Error())
	}

	return conn
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
