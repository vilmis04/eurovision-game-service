package country

import (
	"database/sql"
	"strings"

	"github.com/vilmis04/eurovision-game-service/internal/admin"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(country *Country) (*int64, error) {
	var err error

	var id int64
	err = r.db.QueryRow(`
		INSERT INTO country (name, code, gameType, year, score, isInFinal, artist, song, orderSemi, orderFinal)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id`,
		country.Name, country.Code, country.GameType, country.Year, country.Score, country.IsInFinal,
		country.Artist, country.Song, country.OrderSemi, country.OrderFinal,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

// baseQuery must select from country with `WHERE year=$1`. Only constant
// fragments are appended to it, request values are always bound as args.
func (r *Repo) queryCountries(gameType string, baseQuery string, name string, year string, db *sql.DB) (*sql.Rows, error) {
	var orderBy string
	queryEnd := ""
	if gameType == "final" {
		baseQuery += " AND isInFinal=true"
		orderBy = "orderFinal"
	}
	if strings.Contains(gameType, "semi") {
		orderBy = "orderSemi"
	}
	if orderBy != "" {
		queryEnd = " ORDER BY " + orderBy + " ASC"
	}

	var query string = ""
	if gameType == "" {
		query = baseQuery
	}
	if gameType != "" && admin.GameType(gameType) == admin.GameTypeFinal {
		query = baseQuery + " AND isInFinal=true"
	}

	var rows *sql.Rows
	var err error
	if query != "" {
		rows, err = db.Query(query, year)
		if err != nil {
			return nil, err
		}
	} else {
		var queryParam string
		if name != "" {
			queryParam = name
			query = baseQuery + " AND name=$2" + queryEnd
		} else {
			queryParam = gameType
			query = baseQuery + " AND gameType=$2" + queryEnd
		}
		rows, err = db.Query(query, year, queryParam)
		if err != nil {
			return nil, err
		}
	}

	return rows, nil
}

// Specify game type for countries in that game type.
// Nothing specified will return all countries in the year.
// Specify name to get the specific country in that year
func (r *Repo) GetCountryList(year string, gameType string, name string) (*[]Country, error) {
	var id int
	var countries []Country = []Country{}

	baseQuery := "SELECT * FROM country WHERE year=$1"
	rows, err := r.queryCountries(gameType, baseQuery, name, year, r.db)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		country := Country{}
		err = rows.Scan(&id, &country.Name, &country.Code, &country.Year, &country.GameType, &country.Score, &country.IsInFinal, &country.Artist, &country.Song, &country.OrderSemi, &country.OrderFinal)
		if err != nil {
			return nil, err
		}

		countries = append(countries, country)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &countries, nil
}

// Specify game type for countries in that game type.
// Nothing specified will return all countries in the year.
// Specify name to get the specific country in that year
func (r *Repo) GetCountrySummary(year string, gameType string, name string) (*[]CountrySummary, error) {
	countries := []CountrySummary{}
	baseQuery := "SELECT name, code, artist, song, orderSemi, orderFinal FROM country WHERE year=$1"
	rows, err := r.queryCountries(gameType, baseQuery, name, year, r.db)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		country := CountrySummary{}
		err = rows.Scan(&country.Name, &country.Code, &country.Artist, &country.Song, &country.OrderSemi, &country.OrderFinal)
		if err != nil {
			return nil, err
		}

		countries = append(countries, country)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &countries, nil
}

// UpdateCountry applies the provided fields (nil fields keep their value)
// and returns the number of updated rows.
func (r *Repo) UpdateCountry(req *UpdateCountryRequest, year int, name string) (int64, error) {
	result, err := r.db.Exec(`
		UPDATE country
		SET isInFinal=COALESCE($1, isInFinal),
			score=COALESCE($2, score),
			orderSemi=COALESCE($3, orderSemi),
			orderFinal=COALESCE($4, orderFinal)
		WHERE year=$5 AND name=$6`,
		req.IsInFinal, req.Score, req.OrderSemi, req.OrderFinal, year, name)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// DeleteCountry deletes the country only when exactly one row matches.
// It returns the number of rows that matched, 0 and >1 delete nothing.
func (r *Repo) DeleteCountry(year int, name string) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	res, err := tx.Exec(`DELETE FROM country WHERE year=$1 AND name=$2`, year, name)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rowsAffected != 1 {
		return rowsAffected, nil // deferred rollback discards the delete
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *Repo) GetFinalists(year uint16) ([]Country, error) {
	query := `
				SELECT * FROM "country" 
				WHERE "year"=$1 AND "isinfinal"=true
			`

	rows, err := r.db.Query(query, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	finalists := []Country{}
	for rows.Next() {
		var id int64
		country := Country{}

		err := rows.Scan(&id, &country.Name, &country.Code, &country.Year, &country.GameType, &country.Score, &country.IsInFinal, &country.Artist, &country.Song, &country.OrderSemi, &country.OrderFinal)
		if err != nil {
			return nil, err
		}

		finalists = append(finalists, country)
	}

	return finalists, nil
}
