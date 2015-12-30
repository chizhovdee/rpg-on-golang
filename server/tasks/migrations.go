package tasks

//import (
//	"github.com/chuckpreslar/gofer"
//	"text/template"
//	"path/filepath"
//	"os"
//	"time"
//	"fmt"
//	"errors"
//	"github.com/rubenv/sql-migrate"
//	_ "github.com/go-sql-driver/mysql"
//	"database/sql"
//	"github.com/chizhovdee/rpg/server/config"
//)
//
//func initData() (*migrate.FileMigrationSource, *sql.DB, error){
//	migrations := &migrate.FileMigrationSource{
//		Dir: "server/db/migrations",
//	}
//
//	db, err := config.OpenDb()
//
//	return migrations, db, err
//}
//
//var MigrationUp = gofer.Register(gofer.Task{
//	Namespace:   "migration",
//	Label:       "up",
//	Description: "Migrates the database to the most recent version available",
//	Action: func(arguments ...string) error {
//		migrations, db, err := initData()
//
//		defer db.Close()
//
//		if err != nil {
//			panic(err.Error())
//		}
//
//		n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
//		if err != nil {
//			panic(err.Error())
//		}
//		fmt.Printf("Applied UP %d migrations!\n", n)
//
//		return nil
//	},
//})
//
//var MigrationDown = gofer.Register(gofer.Task{
//	Namespace:   "migration",
//	Label:       "down",
//	Description: "Undo a database migration",
//	Action: func(arguments ...string) error {
//		migrations, db, err := initData()
//
//		defer db.Close()
//
//		if err != nil {
//			panic(err.Error())
//		}
//
//		n, err := migrate.Exec(db, "mysql", migrations, migrate.Down)
//		if err != nil {
//			panic(err.Error())
//		}
//		fmt.Printf("Applied DOWN %d migrations!\n", n)
//
//		return nil
//	},
//})
//
//var MigrationsCreate = gofer.Register(gofer.Task{
//	Namespace:   "migration",
//	Label:       "create",
//	Description: "Create database migration",
//	Action: func(arguments ...string) error {
//		if len(arguments) == 0 {
//			return errors.New("\n No name migration. \n Usage migration:create your_migration_name")
//		}
//
//		dir := filepath.Join("db", "migrations")
//
//		if err := os.MkdirAll(dir, 0777); err != nil {
//			return err
//		}
//
//		t := time.Now()
//
//		timestamp := t.Format("20060102150405")
//		filename := fmt.Sprintf("%v_%v.%v", timestamp, arguments[0], "sql")
//
//		fpath := filepath.Join(dir, filename)
//
//		successPath, err := writeTemplateToFile(fpath, goMigrationTemplate, timestamp)
//
//		if err != nil {
//			return err
//		}
//
//		absolutePath, err := filepath.Abs(successPath)
//
//		if err != nil {
//			return err
//		}
//
//		fmt.Println("migration: created", absolutePath)
//
//		return nil
//	},
//})
//
//var goMigrationTemplate = template.Must(template.New("go-migration-sql").Parse(
//`-- +migrate Up
//-- SQL in section 'Up' is executed when this migration is applied
//
//
//
//-- +migrate Down
//-- SQL section 'Down' is executed when this migration is rolled back
//`))
//
//func writeTemplateToFile(path string, t *template.Template, data interface{}) (string, error) {
//	f, e := os.Create(path)
//	if e != nil {
//		return "", e
//	}
//	defer f.Close()
//
//	e = t.Execute(f, data)
//	if e != nil {
//		return "", e
//	}
//
//	return f.Name(), nil
//}