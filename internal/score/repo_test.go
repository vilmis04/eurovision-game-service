package score

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vilmis04/eurovision-game-service/internal/storage/storagetest"
)

func TestInitializeScoresBindsCountryValues(t *testing.T) {
	st, mock := storagetest.New(t, "score")
	repo := &Repo{storage: st}

	malicious := "x', false, 0); DROP TABLE score;--"
	mock.ExpectQuery(`SELECT name, gametype FROM country WHERE year=$1`).
		WithArgs(2026).
		WillReturnRows(sqlmock.NewRows([]string{"name", "gametype"}).
			AddRow("Serbia", "semi1").
			AddRow(malicious, "final"))
	mock.ExpectBegin()
	insert := `INSERT INTO score ("user", country, year, gametype, inFinal, position) VALUES ($1, $2, $3, $4, false, 0)`
	mock.ExpectPrepare(insert)
	mock.ExpectExec(insert).WithArgs("bob", "Serbia", 2026, "semi1").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(insert).WithArgs("bob", malicious, 2026, "final").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	scores, err := repo.InitializeScores("bob", 2026)
	if err != nil {
		t.Fatal(err)
	}
	if len(scores) != 2 || scores[1].Country != malicious {
		t.Fatalf("unexpected scores %+v", scores)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestInitializeScoresWithoutCountriesInsertsNothing(t *testing.T) {
	st, mock := storagetest.New(t, "score")
	repo := &Repo{storage: st}

	mock.ExpectQuery(`SELECT name, gametype FROM country WHERE year=$1`).
		WithArgs(2026).
		WillReturnRows(sqlmock.NewRows([]string{"name", "gametype"}))

	scores, err := repo.InitializeScores("bob", 2026)
	if err != nil || len(scores) != 0 {
		t.Fatalf("scores = %v, err = %v", scores, err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
