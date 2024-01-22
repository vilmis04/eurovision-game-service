package country

import (
	"fmt"

	"github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/storage"
)

type CountryRepo struct {
	storage *storage.Storage
}

func NewRepo() *CountryRepo {
	return &CountryRepo{
		storage: storage.New("country"),
	}
}

func (r *CountryRepo) Create(country *Country) (*int64, error) {
	db, err := r.storage.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := fmt.Sprintf("INSERT INTO %v (name, gameType, year, score, isInFinal) VALUES ($1, $2, $3, $4, $5)", r.storage.Table)
	result, err := db.Exec(query, country.Name, country.GameType, country.Year, country.Score, country.IsInFinal)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &id, nil
}
