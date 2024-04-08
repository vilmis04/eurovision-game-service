package group

import (
	"database/sql"
	"fmt"
	"html"
	"strings"

	"github.com/lib/pq"
	"github.com/vilmis04/eurovision-game-service/internal/storage"
)

type Repo struct {
	storage.Storage
}

func NewRepo() *Repo {
	return &Repo{
		Storage: *storage.New("group"),
	}
}

// To get all groups of the user, provide "" (empty string) as groupName
func (r *Repo) GetGroupList(owner string, groupName string) (*[]Group, error) {
	db, err := r.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// TODO: update this method (or add a new one) to get all the groups of a member (not owner)
	var rows *sql.Rows
	query := fmt.Sprintf(`
		SELECT * FROM "%v" 
		WHERE owner=$1`, r.Table)
	if groupName != "" {
		query = fmt.Sprintf("%v AND name=$2", query)
		rows, err = db.Query(query, owner, groupName)
	} else {
		rows, err = db.Query(query, owner)
	}
	if err != nil {
		return nil, fmt.Errorf("group query error for %s: %v", owner, err)
	}
	defer rows.Close()

	var groups []Group
	var id int
	for rows.Next() {
		group := Group{}
		var membersResponse string

		err := rows.Scan(&id, &group.Name, &group.Owner, &membersResponse, &group.DateCreated)
		if err != nil {
			return nil, fmt.Errorf("group row scan err: %v", err)
		}

		membersResponse, _ = strings.CutPrefix(membersResponse, "{")
		membersResponse, _ = strings.CutSuffix(membersResponse, "}")
		group.Members = strings.Split(membersResponse, ",")

		groups = append(groups, group)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("group row err: %v", err)
	}

	return &groups, nil
}

func (r *Repo) CreateGroup(group *Group) (*int64, error) {
	db, err := r.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := fmt.Sprintf(`
	INSERT INTO "%v" (name, owner, members, dateCreated)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`, r.Table)

	var id int64
	err = db.QueryRow(query, group.Name, group.Owner, pq.Array(group.Members), group.DateCreated).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (r *Repo) GetGroupNames(owner string) (*([]string), error) {
	db, err := r.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	names := []string{}
	query := fmt.Sprintf(`SELECT name FROM "%v" WHERE owner=$1`, r.Table)
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

		names = append(names, name)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &names, nil
}

func (r *Repo) UpdateMembers(owner string, name string, groupMembers []string) error {
	db, err := r.ConnectToDB()
	if err != nil {
		return fmt.Errorf("conn error: %v", err)
	}
	defer db.Close()

	// TODO: check if quotes around the %v are needed - could it break if owner or group.Name is 'group' (for example)
	query := fmt.Sprintf(`
		UPDATE "%v" 
		SET members=$1 
		WHERE owner='%v' AND name='%v'
	`, r.Table, html.EscapeString(owner), html.EscapeString(name))

	fmt.Println(query)
	_, err = db.Exec(query, pq.Array(groupMembers))
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}

	return nil
}

func (r *Repo) DeleteGroup(owner string, name string) error {
	db, err := r.ConnectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := fmt.Sprintf(`DELETE FROM "%v" WHERE owner=$1 AND name=$2`, r.Table)
	_, err = db.Exec(query, owner, name)
	if err != nil {
		return err
	}

	return nil
}
