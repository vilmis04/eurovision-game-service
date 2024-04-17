package country

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/vilmis04/eurovision-game-service/internal/admin"
	"github.com/vilmis04/eurovision-game-service/internal/storage"
)

type Repo struct {
	storage *storage.Storage
}

func NewRepo() *Repo {
	return &Repo{
		storage: storage.New("country"),
	}
}

func (r *Repo) Create(country *Country) (*int64, error) {
	db, err := r.storage.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := fmt.Sprintf(`
		INSERT INTO %v (name, code, gameType, year, score, isInFinal, orderSemi, orderFinal) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, r.storage.Table)
	result, err := db.Exec(query, country.Name, country.Code, country.GameType, country.Year, country.Score, country.IsInFinal, country.OrderSemi, country.OrderFinal)

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
func (r *Repo) GetCountryList(year string, gameType string, name string) (*[]Country, error) {
	var query string = ""
	var rows *sql.Rows
	var id int
	var countries []Country = []Country{}
	baseQuery := fmt.Sprintf("SELECT * FROM %v WHERE year=$1", r.storage.Table)
	var orderBy string
	queryEnd := ""
	if gameType == "final" {
		orderBy = "orderFinal"
	}
	if strings.Contains(gameType, "semi") {
		orderBy = "orderSemi"
	}
	if orderBy != "" {
		queryEnd = fmt.Sprintf(`
			ORDER BY %v ASC`, orderBy)
	}

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
			query = fmt.Sprintf("%v AND name=$2 %v", baseQuery, queryEnd)
		} else {
			queryParam = gameType
			query = fmt.Sprintf("%v AND gameType=$2 %v", baseQuery, queryEnd)
		}
		rows, err = db.Query(query, year, queryParam)
		if err != nil {
			return nil, err
		}

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

func (r *Repo) UpdateCountry(req *UpdateCountryRequest, params *map[string]string) error {
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

func (r *Repo) DeleteCountry(year string, name string) error {
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
