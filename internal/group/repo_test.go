package group

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vilmis04/eurovision-game-service/internal/storage/storagetest"
)

func TestRepoUpdateMembersIsParameterised(t *testing.T) {
	st, mock := storagetest.New(t, "group")
	repo := &Repo{Storage: *st}

	mock.ExpectExec(`UPDATE "group" SET members=$1 WHERE id=$2`).
		WithArgs(sqlmock.AnyArg(), int64(7)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.UpdateMembers(7, []string{"alice", "bob'); DROP TABLE \"group\";--"})
	if err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestRepoDeleteGroupIsParameterised(t *testing.T) {
	st, mock := storagetest.New(t, "group")
	repo := &Repo{Storage: *st}

	mock.ExpectExec(`DELETE FROM "group" WHERE owner=$1 AND id=$2`).
		WithArgs("o' OR '1'='1", int64(3)).
		WillReturnResult(sqlmock.NewResult(0, 0))

	if err := repo.DeleteGroup("o' OR '1'='1", 3); err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
