package group

import (
	"fmt"

	"github.com/vilmis04/eurovision-game-service/internal/storage"
)

type Repo struct {
	storage *storage.Storage
}

func NewRepo() *Repo {
	return &Repo{
		storage: storage.New("group"),
	}
}

func (r *Repo) GetGroupList(owner string, groupName string) (*[]Group, error) {
	db, err := r.storage.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT * FROM group WHERE owner=$1"
	if groupName != "" {
		query = fmt.Sprintf("%v AND name=$2", query)
	}

	rows, err := db.Query(query, owner, groupName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []Group
	var id int
	for rows.Next() {
		group := Group{}

		err := rows.Scan(&id, &group.Name, &group.Members, &group.Owner, &group.DateCreated)
		if err != nil {
			return nil, err
		}

		groups = append(groups, group)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return &groups, nil
}
