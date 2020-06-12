package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"net/url"
	"os"
	"strings"
	"time"
)

func MigrateVersion(driver string, dsn string, dir string) (uint, bool, error) {
	mig, err := NewMigration(driver, dsn, dir)

	if err != nil {
		return 0, false, err
	}

	return mig.Version()
}

func MigrateStep(driver string, dsn string, dir string, step int) error {
	mig, err := NewMigration(driver, dsn, dir)

	if err != nil {
		return err
	}

	if err := mig.Steps(step); err != nil {
		return err
	}

	return nil
}

func MigrateTo(driver string, dsn string, dir string, version uint) error {
	mig, err := NewMigration(driver, dsn, dir)

	if err != nil {
		return err
	}

	if err := mig.Migrate(version); err != nil {
		return err
	}

	return nil
}

func MigrateUp(driver string, dsn string, dir string) error {
	mig, err := NewMigration(driver, dsn, dir)

	if err != nil {
		return err
	}

	if err := mig.Up(); err != nil {
		return err
	}

	return nil
}

func MigrateDown(driver string, dsn string, dir string) error {
	mig, err := NewMigration(driver, dsn, dir)

	if err != nil {
		return err
	}

	if err := mig.Down(); err != nil {
		return err
	}

	return nil
}

func writeMigrationFile(filename string) {
	file, err := os.Create(filename)

	if err != nil {
		file.Close()
		fmt.Printf("Cannot create file, %s", err)
		return
	}

	d := []string{
		"BEGIN;",
		"",
		"-- Write here your migration",
		"",
		"COMMIT;",
	}

	for _, v := range d {
		_, err = fmt.Fprintln(file, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if err = file.Close(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("file written successfully", filename)
}

func NewMigrationFile(dir string, commit string) {
	timestamp := time.Now().Format("2006-01-02 15.04.05")

	filenameT := strings.ReplaceAll(timestamp, " ", "")
	filenameT = strings.ReplaceAll(filenameT, "-", "")
	filenameT = strings.ReplaceAll(filenameT, ".", "")

	filenameM := commit
	filenameM = strings.ReplaceAll(filenameM, " ", "_")
	filenameM = strings.ToLower(filenameM)

	filenameU := fmt.Sprintf("%s/%s_%s.up.sql", dir, filenameT, filenameM)
	filenameD := fmt.Sprintf("%s/%s_%s.down.sql", dir, filenameT, filenameM)

	fmt.Println(filenameU)
	fmt.Println(filenameD)

	writeMigrationFile(filenameU)
	writeMigrationFile(filenameD)
}

func NewMigration(driver string, dsn string, dir string) (*migrate.Migrate, error) {
	db, err := NewMigrateConnection(driver, dsn)

	if err != nil {
		return nil, err
	}

	dr, err := NewMigrateDriver(db)

	if err != nil {
		return nil, err
	}

	ur, err := url.Parse(dsn)

	if err != nil {
		return nil, err
	}

	return NewMigrateDir(dr, dir, strings.TrimPrefix(ur.Path, "/"))
}

func NewMigrateDir(driver database.Driver, dir string, name string) (*migrate.Migrate, error) {
	file := fmt.Sprintf("file://%s", dir)
	return migrate.NewWithDatabaseInstance(
		file,
		name,
		driver,
	)
}

func NewMigrateDriver(db *sql.DB) (database.Driver, error) {
	return postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: "migrate_version",
	})
}

func NewMigrateConnection(driver string, dsn string) (*sql.DB, error) {
	return sql.Open(driver, dsn)
}
