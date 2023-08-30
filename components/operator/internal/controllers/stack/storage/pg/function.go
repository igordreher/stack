package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/formancehq/operator/apis/stack/v1beta3"
)

var (
	ErrNotExisting = errors.New("not existing")
)

func Exists(db *sql.DB, databaseName string) error {
	res, err := db.Exec("SELECT datname FROM pg_catalog.pg_database WHERE lower(datname) = lower('" + databaseName + "');")
	if err != nil {
		return err
	}

	l, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if l == 0 {
		return ErrNotExisting
	}

	return nil
}

func DropDB(db *sql.DB, stackName string, serviceName string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + fmt.Sprintf("\"%s-%s\"", stackName, serviceName))
	if err != nil {
		return err
	}

	return nil
}

func OpenClient(config v1beta3.PostgresConfig) (*sql.DB, error) {
	debug := true
	return OpenSQLDB(ConnectionOptions{
		DatabaseSourceName: config.DSN(),
		Debug:              debug,
		Trace:              debug,
		Writer:             os.Stdout,
		MaxIdleConns:       20,
		ConnMaxIdleTime:    time.Minute,
		MaxOpenConns:       20,
	})
}
