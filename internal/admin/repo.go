package admin

import (
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

	row := db.QueryRow(`SELECT * FROM admin_config WHERE id=1`)
	err = row.Scan(&id, &config.Year, &config.GameType, &config.IsVotingAcitve, &config.VotingEnd)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (r *Repo) UpdateConfig(body *adminConfigRequestBody) error {
	db, err := r.storage.ConnectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if body.GameType != nil {
		_, err = db.Exec(`UPDATE admin_config SET gameType=$1 WHERE id=1`, *body.GameType)
		if err != nil {
			return err
		}
	}
	if body.Year != nil {
		_, err = db.Exec(`UPDATE admin_config SET year=$1 WHERE id=1`, *body.Year)
		if err != nil {
			return err
		}
	}
	if body.IsVotingAcitve != nil {
		_, err = db.Exec(`UPDATE admin_config SET isVotingActive=$1 WHERE id=1`, *body.IsVotingAcitve)
		if err != nil {
			return err
		}
	}

	return nil
}
