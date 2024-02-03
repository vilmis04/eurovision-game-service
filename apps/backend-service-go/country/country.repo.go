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
// Nothing specified will return all countries in the year.
// Specify name to get the specific country in that year
func (r *CountryRepo) GetCountryList(year string, gameType string, name string) (*[]Country, error) {
	var query string = ""
	var rows *sql.Rows
	var id int
	var countries []Country = []Country{}
	baseQuery := fmt.Sprintf("SELECT * FROM %v WHERE year=$1", r.storage.Table)

	db, err := r.storage.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if gameType == "" {
		query = baseQuery
	}
	if gameType != "" && admin.GameType(gameType) == admin.GameTypeFinal {
		query = fmt.Sprintf("%v AND isInFinal=true", baseQuery)
	}
	if query != "" {
		rows, err = db.Query(query, year)
		if err != nil {
			return nil, err
		}
	} else {
		var queryParam string
		if name != "" {
			queryParam = name
			query = fmt.Sprintf("%v AND name=$2", baseQuery)
		} else {
			queryParam = gameType
			query = fmt.Sprintf("%v AND gameType=$2", baseQuery)
		}
		rows, err = db.Query(query, year, queryParam)
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

func (r *CountryRepo) UpdateCountry(req *UpdateCountryRequest, params *map[string]string) error {
	db, err := r.storage.ConnectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var formatQuery = func(queryParam string) string {
		return fmt.Sprintf("UPDATE %v SET %v=$1 WHERE year=%v AND name=%v", r.storage.Table, queryParam, (*params)["year"], (*params)["name"])
	}

	if req.IsInFinal != nil {
		_, err := db.Exec(formatQuery("isInFinal"), *req.IsInFinal)
		if err != nil {
			return err
		}
	}

	if req.Score != nil {
		_, err := db.Exec(formatQuery("Score"), *req.Score)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *CountryRepo) DeleteCountry(year string, name string) error {
	db, err := r.storage.ConnectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() error {
		if err != nil {
			tx.Rollback()
			return err
		}
		return nil
	}()

	res, err := tx.Exec(fmt.Sprintf(`
		DELETE %v
		WHERE year=$1 AND name=$2`, r.storage.Table), year, name)

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected != 1 {
		tx.Rollback()
		return fmt.Errorf("multiple or none rows affected")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
