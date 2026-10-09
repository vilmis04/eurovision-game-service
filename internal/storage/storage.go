package storage

import (
	"database/sql"
	"fmt"
	"os"
)

type Storage struct {
	ConnString string
	Table      string
	// Driver is the database/sql driver name, "postgres" unless overridden (e.g. in tests).
	Driver string
}

func New(table string) *Storage {
	return &Storage{
		ConnString: fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB")),
		Table:  table,
		Driver: "postgres",
	}
}

func (s *Storage) ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open(s.Driver, s.ConnString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
