package country

import (
	"database/sql"
	"fmt"

	"github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/admin"
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

// Specify game type for countries in that game type.
// Nothing specified will return all countries in the year
func (r *CountryRepo) GetAll(year uint16, gameType *admin.GameType) (*[]Country, error) {
	var query string = ""
	var rows *sql.Rows
	var id int
	var countries []Country = []Country{}

	db, err := r.storage.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if gameType == nil {
		query = fmt.Sprintf("SELECT * FROM %v WHERE year=$1", r.storage.Table)
	}
	if gameType != nil && *gameType == admin.GameTypeFinal {
		query = fmt.Sprintf("SELECT * FROM %v WHERE year=$1 AND isInFinal=true", r.storage.Table)
	}
	if query != "" {
		rows, err = db.Query(query, year)
		if err != nil {
			return nil, err
		}
	} else {
		query = fmt.Sprintf("SELECT * FROM %v WHERE year=$1 AND gameType=$2", r.storage.Table)
		rows, err = db.Query(query, year, *gameType)
		if err != nil {
			return nil, err
		}

	}
	defer rows.Close()

	for rows.Next() {
		country := Country{}

		err = rows.Scan(&id, &country.Name, &country.Year, &country.GameType, &country.Score, &country.IsInFinal, &country.Artist, &country.Song)
		if err != nil {
			return nil, err
		}

		countries = append(countries, country)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return &countries, nil
}
