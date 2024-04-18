package score

import (
	"database/sql"
	"log"

	"github.com/vilmis04/eurovision-game-service/internal/storage"
)

type Repo struct {
	storage *storage.Storage
}

func NewRepo() *Repo {
	return &Repo{
		storage: storage.New("score"),
	}
}

func (r *Repo) GetScore(user string, country string, year uint16) (*Score, error) {
	db, err := r.storage.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	score := Score{}
	query := "SELECT * FROM score WHERE (user=$1 AND country=$2 AND year=$3)"
	row := db.QueryRow(query, user, country, year)

	var id int64
	err = row.Scan(id, score.Country, score.Year, score.User, score.InFinal, score.Points)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("no rows found for %v of %v in %v\n", user, country, year)
			return nil, nil
		}
		return nil, err
	}

	return &score, nil
}

// TODO: update GetAllScores to return list of all scores
func (r *Repo) GetAllScores(user string, country string, year uint16) (*Score, error) {
	db, err := r.storage.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	score := Score{}
	query := "SELECT * FROM score WHERE (user=$1 AND country=$2 AND year=$3)"
	row := db.QueryRow(query, user, country, year)

	var id int64
	err = row.Scan(id, score.Country, score.Year, score.User, score.InFinal, score.Points)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("no rows found for %v of %v in %v\n", user, country, year)
			return nil, nil
		}
		return nil, err
	}

	return &score, nil
}
