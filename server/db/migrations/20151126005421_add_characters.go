
package main

import (
	"database/sql"
	"fmt"
)

// Up is executed when this migration is applied
func Up_20151126005421(txn *sql.Tx) {
	sqlQuery := `
		create table characters (
		  id int(11) NOT NULL AUTO_INCREMENT,
          PRIMARY KEY(id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci`

	_, err := txn.Exec(sqlQuery)

	if err != nil {
	    txn.Rollback()

	    panic(err.Error())
	}

	fmt.Println("Миграция завершена успешно")
}

// Down is executed when this migration is rolled back
func Down_20151126005421(txn *sql.Tx) {
	_, err := txn.Exec("drop table characters")

	if err != nil {
		txn.Rollback()

		panic(err.Error())
	}

	fmt.Println("Миграция завершена успешно")
}
