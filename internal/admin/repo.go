package admin

import (
	"fmt"

	"github.com/vilmis04/eurovision-game-service/internal/storage"
)

type Repo struct {
	storage *storage.Storage
}

func NewRepo() *Repo {
	return &Repo{
		storage: storage.New("admin_config"),
	}
}

func (r *Repo) GetConfig() (*Admin, error) {
	db, err := r.storage.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var config Admin
	var id uint16

	row := db.QueryRow(fmt.Sprintf("SELECT * FROM %v WHERE id=1", r.storage.Table))
	err = row.Scan(&id, &config.Year, &config.GameType, &config.IsVotingAcitve, &config.VotingEnd)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (r *Repo) formatConfigUpdateQuery(property string) string {
	return fmt.Sprintf("UPDATE %v SET %v=$1 WHERE id=1", r.storage.Table, property)
}

func (r *Repo) UpdateConfig(body *adminConfigRequestBody) error {
	db, err := r.storage.ConnectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if body.GameType != nil {
		_, err = db.Exec(r.formatConfigUpdateQuery("gameType"), *body.GameType)
		if err != nil {
			return err
		}
	}
	if body.Year != nil {
		_, err = db.Exec(r.formatConfigUpdateQuery("year"), *body.Year)
		if err != nil {
			return err
		}
	}
	if body.IsVotingAcitve != nil {
		_, err = db.Exec(r.formatConfigUpdateQuery("isVotingActive"), *body.IsVotingAcitve)
		if err != nil {
			return err
		}
	}

	return nil
}
