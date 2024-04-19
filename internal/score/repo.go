package score

import (
	"database/sql"
	"log"

	"github.com/vilmis04/eurovision-game-service/internal/admin"
	"github.com/vilmis04/eurovision-game-service/internal/country"
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

func (r *Repo) GetAllScores(user string, gameType admin.GameType, year uint16) ([]ScoreResponse, error) {

	db, err := r.storage.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	scores := []ScoreResponse{}
	query := `SELECT country, infinal, points FROM score WHERE ("user"=$1 AND "year"=$2 AND "gametype"=$3)`
	rows, err := db.Query(query, user, year, gameType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		score := ScoreResponse{}
		err = rows.Scan(&score.Country, &score.InFinal, &score.Points)
		if err != nil {
			return nil, err
		}

		scores = append(scores, score)
	}
	log.Printf("found %v scores for %v in %v\n", len(scores), user, year)

	return scores, nil
}

func (r *Repo) InitializeScores(user string, year uint16, gameType admin.GameType, countries *[]country.Country) ([]ScoreResponse, error) {
	db, err := r.storage.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	scores := []ScoreResponse{}
	for _, country := range *countries {
		score := Score{
			Country:  country.Name,
			Year:     year,
			User:     user,
			GameType: gameType,
			InFinal:  false,
			Points:   0,
		}

		query := `INSERT INTO score ("user", country, year, gametype, inFinal, points) VALUES ($1, $2, $3, $4, $5, $6)`
		_, err = db.Exec(query, score.User, score.Country, score.Year, score.GameType, score.InFinal, score.Points)
		if err != nil {
			return nil, err
		}

		scores = append(scores, ScoreResponse{
			Country: score.Country,
			InFinal: score.InFinal,
			Points:  score.Points,
		})
	}

	return scores, nil
}
