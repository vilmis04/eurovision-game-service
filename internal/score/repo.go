package score

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/vilmis04/eurovision-game-service/internal/admin"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) GetScore(user string, country string, year uint16) (*Score, error) {
	var err error

	score := Score{}
	query := `SELECT * FROM score WHERE ("user"=$1 AND "country"=$2 AND "year"=$3)`
	row := r.db.QueryRow(query, user, country, year)

	var id int64
	err = row.Scan(&id, &score.Country, &score.Year, &score.GameType, &score.User, &score.InFinal, &score.Position)
	if err != nil {
		return nil, err
	}

	return &score, nil
}

func (r *Repo) GetAllOrSemiScores(user string, gameType admin.GameType, year uint16) ([]ScoreResponse, error) {
	var err error

	scores := []ScoreResponse{}
	var rows *sql.Rows
	if gameType == "" {
		rows, err = r.db.Query(`SELECT country, infinal, position FROM score WHERE ("user"=$1 AND "year"=$2)`, user, year)
	} else {
		rows, err = r.db.Query(`SELECT country, infinal, position FROM score WHERE ("user"=$1 AND "year"=$2 AND "gametype"=$3)`, user, year, gameType)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		score := ScoreResponse{}
		err = rows.Scan(&score.Country, &score.InFinal, &score.Position)
		if err != nil {
			return nil, err
		}

		scores = append(scores, score)
	}
	log.Printf("found %v scores for %v in %v\n", len(scores), user, year)

	return scores, nil
}

// EnsureScores inserts a row for every country of the year that the user does not
// have yet, in one statement. The unique key on (user, country, year) makes it
// idempotent, so repeated calls never create duplicates.
func (r *Repo) EnsureScores(user string, year uint16) error {
	_, err := r.db.Exec(`
		INSERT INTO score ("user", country, year, gametype, infinal, position)
		SELECT $1, c.name, c.year, c.gametype, false, 0
		FROM country c
		WHERE c.year = $2
		ON CONFLICT ("user", country, year) DO NOTHING`, user, year)
	if err != nil {
		return fmt.Errorf("insert scores err: %v", err)
	}

	return nil
}

func (r *Repo) UpdateScore(user string, score *Score) error {
	query := `UPDATE score
			  SET infinal=$1, position=$2
			  WHERE ("user"=$3 AND country=$4 AND year=$5)`
	result, err := r.db.Exec(query, score.InFinal, score.Position, score.User, score.Country, score.Year)
	if err != nil {
		return fmt.Errorf("failed to update score for %v of %v in %v: %v", user, score.Country, score.Year, err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated (user: %s, country: %s, year: %d)", score.User, score.Country, score.Year)
	}

	return nil
}

func (r *Repo) GetFinalScores(user string, year uint16) ([]ScoreResponse, error) {
	query := `
				SELECT s.country, s.infinal, s.position
				FROM country c
				JOIN score s ON c.name = s.country
				WHERE c.isinfinal = true AND s.user=$1 AND c.year=$2 AND s.year=$2
			`
	rows, err := r.db.Query(query, user, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	scores := []ScoreResponse{}
	for rows.Next() {
		score := ScoreResponse{}
		err = rows.Scan(&score.Country, &score.InFinal, &score.Position)
		if err != nil {
			return nil, err
		}

		scores = append(scores, score)
	}

	return scores, nil
}

func (r *Repo) GetMultipleScores(userList []string, year uint16) (map[string][]Score, error) {
	scores := make(map[string][]Score)
	if len(userList) == 0 {
		return scores, nil
	}

	// only positional placeholders are generated here, user values are bound as args
	var queryBuilder strings.Builder
	queryBuilder.WriteString(`SELECT * FROM score WHERE year=$1 AND "user" IN (`)
	for i := range userList {
		queryBuilder.WriteString("$" + strconv.Itoa(i+2))
		if i != len(userList)-1 {
			queryBuilder.WriteString(",")
		}
	}
	queryBuilder.WriteString(")")
	query := queryBuilder.String()

	args := make([]interface{}, len(userList)+1) // adding year as an arg
	args[0] = year
	for i, v := range userList {
		args[i+1] = v // offset the year
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		score := Score{}
		err = rows.Scan(&id, &score.Country, &score.Year, &score.GameType, &score.User, &score.InFinal, &score.Position)
		if err != nil {
			return nil, err
		}

		if _, ok := scores[score.User]; !ok {
			scores[score.User] = []Score{score}
			continue
		}

		scores[score.User] = append(scores[score.User], score)
	}

	return scores, nil
}
