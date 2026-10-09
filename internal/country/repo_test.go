package country

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vilmis04/eurovision-game-service/internal/storage/storagetest"
)

const maliciousName = "x'; DROP TABLE country;--"

func TestRepoUpdateCountryIsParameterised(t *testing.T) {
	st, mock := storagetest.New(t, "country")
	repo := &Repo{storage: st}

	inFinal := true
	mock.ExpectExec(`UPDATE country SET isInFinal=COALESCE($1, isInFinal), score=COALESCE($2, score), orderSemi=COALESCE($3, orderSemi), orderFinal=COALESCE($4, orderFinal) WHERE year=$5 AND name=$6`).
		WithArgs(true, nil, nil, nil, 2026, maliciousName).
		WillReturnResult(sqlmock.NewResult(0, 1))

	updated, err := repo.UpdateCountry(&UpdateCountryRequest{IsInFinal: &inFinal}, 2026, maliciousName)
	if err != nil {
		t.Fatal(err)
	}
	if updated != 1 {
		t.Fatalf("updated = %d, want 1", updated)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestRepoDeleteCountryCommitsOnlyForExactlyOneRow(t *testing.T) {
	t.Run("one row commits", func(t *testing.T) {
		st, mock := storagetest.New(t, "country")
		repo := &Repo{storage: st}

		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM country WHERE year=$1 AND name=$2`).
			WithArgs(2026, maliciousName).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		deleted, err := repo.DeleteCountry(2026, maliciousName)
		if err != nil || deleted != 1 {
			t.Fatalf("deleted = %d, err = %v", deleted, err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("several rows roll back", func(t *testing.T) {
		st, mock := storagetest.New(t, "country")
		repo := &Repo{storage: st}

		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM country WHERE year=$1 AND name=$2`).
			WithArgs(2026, "Serbia").
			WillReturnResult(sqlmock.NewResult(0, 2))
		mock.ExpectRollback()

		deleted, err := repo.DeleteCountry(2026, "Serbia")
		if err != nil || deleted != 2 {
			t.Fatalf("deleted = %d, err = %v", deleted, err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatal(err)
		}
	})
}
