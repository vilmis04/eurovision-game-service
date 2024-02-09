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

func (r *Repo) CreateGroup(group *Group) (*int64, error) {
	db, err := r.storage.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := fmt.Sprintf(`
	INSERT INTO %v (name, owner, members, dateCreated)
	VALUES ($1, $2, $3, $4)
	`, r.storage.Table)
	result, err := db.Exec(query, group.Name, group.Owner, group.Members, group.DateCreated)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (r *Repo) GetGroupNames(owner string) (*(map[string]string), error) {
	db, err := r.storage.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	names := map[string]string{}
	query := fmt.Sprintf("SELECT name FROM %v WHERE owner=$1", r.storage.Table)
	rows, err := db.Query(query, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}

		names[name] = name
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &names, nil
}
