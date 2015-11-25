package tasks

import (
	"github.com/chuckpreslar/gofer"
	"text/template"
	"path/filepath"
	"os"
	"time"
	"fmt"
	"errors"
)

var MigrationsCreate = gofer.Register(gofer.Task{
	Namespace:   "migration",
	Label:       "create",
	Description: "Create database migration",
	Action: func(arguments ...string) error {
		if len(arguments) == 0 {
			return errors.New("\n No name migration. \n Usage migration:create your_migration_name")
		}

		dir := filepath.Join("db", "migrations")

		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}

		t := time.Now()

		timestamp := t.Format("20060102150405")
		filename := fmt.Sprintf("%v_%v.%v", timestamp, arguments[0], "go")

		fpath := filepath.Join(dir, filename)

		successPath, err := writeTemplateToFile(fpath, goMigrationTemplate, timestamp)

		if err != nil {
			return err
		}

		absolutePath, err := filepath.Abs(successPath)

		if err != nil {
			return err
		}

		fmt.Println("migration: created", absolutePath)

		return nil
	},
})

var goMigrationTemplate = template.Must(template.New("goose.go-migration").Parse(`
package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_{{ . }}(txn *sql.Tx) {
	_, err := txn.Exec("Your query")

	if err != nil {
	    txn.Rollback()

	    log.Println(err.Error())
	}

	fmt.Println("Миграция завершена успешно!")
}

// Down is executed when this migration is rolled back
func Down_{{ . }}(txn *sql.Tx) {
    _, err := txn.Exec("Your query")

	if err != nil {
	    txn.Rollback()

	    log.Println(err.Error())
	}

	fmt.Println("Миграция завершена успешно!")
}
`))

func writeTemplateToFile(path string, t *template.Template, data interface{}) (string, error) {
	f, e := os.Create(path)
	if e != nil {
		return "", e
	}
	defer f.Close()

	e = t.Execute(f, data)
	if e != nil {
		return "", e
	}

	return f.Name(), nil
}