package storage

import (
	"database/sql"
	"fmt"
	"os"
)

type Storage struct {
	ConnString string
	Table      string
}

func New(table string) *Storage {
	return &Storage{
		ConnString: fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB")),
		Table: table,
	}
}

func (s *Storage) ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", s.ConnString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
