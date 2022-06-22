package postgress

import (
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

//go:embed migrations/*.sql
var fs embed.FS

type PostgresDB struct {
	log logrus.FieldLogger

	pool *pgxpool.Pool

	connstr string
}

func New(log logrus.FieldLogger, connStr string) *PostgresDB {
	return &PostgresDB{
		log:     log,
		connstr: connStr,
	}
}

func (db *PostgresDB) RunMigrations() error {
	d, err := iofs.New(fs, "migrations")
	if err != nil {
		return fmt.Errorf("failed to create vfs: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, db.connstr)
	if err != nil {
		return fmt.Errorf("failed to initialize: %w", err)
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return err
		} else {
			return fmt.Errorf("golang-migrate failed to run Up(): %w", err)
		}
	}
	fmt.Println("ran migrations successfully")

	return nil
}
