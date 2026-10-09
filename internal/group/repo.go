package group

import (
	"database/sql"
	"errors"
	"fmt"

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

const groupColumns = `id, name, owner, members, datecreated`

func scanGroup(row interface{ Scan(dest ...any) error }) (*Group, error) {
	group := Group{}
	err := row.Scan(&group.Id, &group.Name, &group.Owner, pq.Array(&group.Members), &group.DateCreated)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

// To get all groups of the user, provide "" (empty string) as groupId.
// groupId must be a validated numeric id.
func (r *Repo) GetGroupList(user string, groupId string) (*[]Group, error) {
	db, err := r.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var rows *sql.Rows
	if groupId != "" {
		rows, err = db.Query(`SELECT `+groupColumns+` FROM "group" WHERE $1 = ANY(members) AND id=$2`, user, groupId)
	} else {
		rows, err = db.Query(`SELECT `+groupColumns+` FROM "group" WHERE $1 = ANY(members)`, user)
	}
	if err != nil {
		return nil, fmt.Errorf("group query error for %s: %v", user, err)
	}
	defer rows.Close()

	var groups []Group
	for rows.Next() {
		group, err := scanGroup(rows)
		if err != nil {
			return nil, fmt.Errorf("group row scan err: %v", err)
		}

		groups = append(groups, *group)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("group row err: %v", err)
	}

	return &groups, nil
}

// GetGroupById returns nil, nil when no such group exists.
func (r *Repo) GetGroupById(id int64) (*Group, error) {
	db, err := r.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	group, err := scanGroup(db.QueryRow(`SELECT `+groupColumns+` FROM "group" WHERE id=$1`, id))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("group by id query error: %v", err)
	}

	return group, nil
}

func (r *Repo) CreateGroup(group *Group) (*int64, error) {
	db, err := r.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var id int64
	err = db.QueryRow(
		`INSERT INTO "group" (name, owner, members, dateCreated) VALUES ($1, $2, $3, $4) RETURNING id`,
		group.Name, group.Owner, pq.Array(group.Members), group.DateCreated,
	).Scan(&id)
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
	rows, err := db.Query(`SELECT name FROM "group" WHERE owner=$1`, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}

		names = append(names, name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &names, nil
}

func (r *Repo) UpdateMembers(id int64, groupMembers []string) error {
	db, err := r.ConnectToDB()
	if err != nil {
		return fmt.Errorf("conn error: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`UPDATE "group" SET members=$1 WHERE id=$2`, pq.Array(groupMembers), id)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}

	return nil
}

func (r *Repo) DeleteGroup(owner string, id int64) error {
	db, err := r.ConnectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`DELETE FROM "group" WHERE owner=$1 AND id=$2`, owner, id)
	if err != nil {
		return err
	}

	return nil
}
