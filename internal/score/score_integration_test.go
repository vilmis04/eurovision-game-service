package score_test

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/vilmis04/eurovision-game-service/internal/admin"
	"github.com/vilmis04/eurovision-game-service/internal/country"
	"github.com/vilmis04/eurovision-game-service/internal/dbtest"
	"github.com/vilmis04/eurovision-game-service/internal/score"
)

// Needs a real Postgres: set TEST_DATABASE_URL to a superuser URL.

func newService(database *sql.DB) *score.Service {
	adminService := admin.NewService(database)

	return score.NewService(database, adminService, country.NewService(database, adminService))
}

func rows(t *testing.T, database *sql.DB, user string) int {
	t.Helper()
	var n int
	if err := database.QueryRow(`SELECT count(*) FROM score WHERE "user" = $1`, user).Scan(&n); err != nil {
		t.Fatal(err)
	}

	return n
}

func TestRepeatedGetNeverCreatesDuplicateRows(t *testing.T) {
	database := dbtest.NewMigrated(t)
	service := newService(database)

	for _, stage := range []string{"semi1", "semi2", "final"} {
		if _, err := database.Exec(`UPDATE admin_config SET gameType = $1 WHERE id = 1`, stage); err != nil {
			t.Fatal(err)
		}
		if _, err := database.Exec(`UPDATE country SET isinfinal = (gametype = 'final') WHERE year = 2026`); err != nil {
			t.Fatal(err)
		}

		var lastBody string
		for i := 0; i < 4; i++ {
			body, err := service.GetAllScores("amy", false)
			if err != nil {
				t.Fatalf("%s call %d: %v", stage, i+1, err)
			}
			lastBody = string(*body)

			if got := rows(t, database, "amy"); got != 34 {
				t.Fatalf("%s call %d: amy has %d score rows, want 34 (one per 2026 country)", stage, i+1, got)
			}
		}

		var response []score.ScoreResponse
		if err := json.Unmarshal([]byte(lastBody), &response); err != nil {
			t.Fatal(err)
		}
		seen := map[string]bool{}
		for _, item := range response {
			if seen[item.Country] {
				t.Fatalf("%s: %s appears twice in the response", stage, item.Country)
			}
			seen[item.Country] = true
		}
		if len(response) == 0 {
			t.Fatalf("%s: empty response", stage)
		}
	}
}

func TestEnsureScoresKeepsExistingPicks(t *testing.T) {
	database := dbtest.NewMigrated(t)
	service := newService(database)

	if err := service.EnsureScores("amy", 2026); err != nil {
		t.Fatal(err)
	}
	if _, err := database.Exec(`UPDATE score SET infinal = true, position = 5 WHERE "user" = 'amy' AND country = 'France'`); err != nil {
		t.Fatal(err)
	}
	if err := service.EnsureScores("amy", 2026); err != nil {
		t.Fatal(err)
	}

	var inFinal bool
	var position int
	if err := database.QueryRow(`SELECT infinal, position FROM score WHERE "user" = 'amy' AND country = 'France'`).Scan(&inFinal, &position); err != nil {
		t.Fatal(err)
	}
	if !inFinal || position != 5 {
		t.Fatalf("existing pick was reset: infinal=%v position=%d", inFinal, position)
	}
}

func TestSecondUserGetsTheirOwnRows(t *testing.T) {
	database := dbtest.NewMigrated(t)
	service := newService(database)

	for _, user := range []string{"amy", "bob", "amy", "bob"} {
		if err := service.EnsureScores(user, 2026); err != nil {
			t.Fatal(err)
		}
	}
	if rows(t, database, "amy") != 34 || rows(t, database, "bob") != 34 {
		t.Fatal("each user must have exactly one row per country")
	}
}
