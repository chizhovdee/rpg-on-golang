package main
import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"log"
	"database/sql"
	"fmt"
)

func existsPendingMigrations(db *sql.DB) bool {
	conf, err := goose.NewDBConf("db", "development", "")
	if err != nil {
		log.Fatal(err)
	}

	// collect all migrations
	min := int64(0)
	max := int64((1 << 63) - 1)
	migrations, e := goose.CollectMigrations(conf.MigrationsDir, min, max)

	if e != nil {
		log.Fatal(e)
	}

	// must ensure that the version table exists if we're running on a pristine DB
	if _, e := goose.EnsureDBVersion(conf, db); e != nil {
		log.Fatal(e)
	}

	for _, m := range migrations {
		row := goose.MigrationRecord{}

		q := fmt.Sprintf(
			"SELECT tstamp, is_applied FROM goose_db_version WHERE version_id=%d ORDER BY tstamp DESC LIMIT 1",
			m.Version,
		)
		e := db.QueryRow(q).Scan(&row.TStamp, &row.IsApplied)

		if e != nil && e != sql.ErrNoRows {
			log.Fatal(e)
		}

		if !row.IsApplied {
			return true
		}
	}

	return false
}