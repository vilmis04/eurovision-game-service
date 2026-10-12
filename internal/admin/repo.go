package admin

import (
	"database/sql"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) GetConfig() (*Admin, error) {
	var config Admin
	var id uint16

	row := r.db.QueryRow(`SELECT * FROM admin_config WHERE id=1`)
	err := row.Scan(&id, &config.Year, &config.GameType, &config.IsVotingAcitve, &config.VotingEnd)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (r *Repo) UpdateConfig(body *adminConfigRequestBody) error {
	var err error

	if body.GameType != nil {
		_, err = r.db.Exec(`UPDATE admin_config SET gameType=$1 WHERE id=1`, *body.GameType)
		if err != nil {
			return err
		}
	}
	if body.Year != nil {
		_, err = r.db.Exec(`UPDATE admin_config SET year=$1 WHERE id=1`, *body.Year)
		if err != nil {
			return err
		}
	}
	if body.IsVotingAcitve != nil {
		_, err = r.db.Exec(`UPDATE admin_config SET isVotingActive=$1 WHERE id=1`, *body.IsVotingAcitve)
		if err != nil {
			return err
		}
	}

	return nil
}
