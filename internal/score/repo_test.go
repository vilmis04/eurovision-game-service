package score

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vilmis04/eurovision-game-service/internal/dbtest"
)

func TestEnsureScoresIsOneIdempotentStatement(t *testing.T) {
	database, mock := dbtest.NewMock(t)
	repo := NewRepo(database)

	mock.ExpectExec(`INSERT INTO score ("user", country, year, gametype, infinal, position) SELECT $1, c.name, c.year, c.gametype, false, 0 FROM country c WHERE c.year = $2 ON CONFLICT ("user", country, year) DO NOTHING`).
		WithArgs("bob'; DROP TABLE score;--", 2026).
		WillReturnResult(sqlmock.NewResult(0, 0))

	if err := repo.EnsureScores("bob'; DROP TABLE score;--", 2026); err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
