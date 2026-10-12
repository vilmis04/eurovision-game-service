package migrations_test

import (
	"context"
	"database/sql"
	"sync"
	"testing"

	"github.com/vilmis04/eurovision-game-service/internal/dbtest"
	"github.com/vilmis04/eurovision-game-service/internal/migrations"
)

// These tests need a real Postgres: set TEST_DATABASE_URL to a superuser URL.

func count(t *testing.T, database *sql.DB, query string, args ...any) int {
	t.Helper()
	var n int
	if err := database.QueryRow(query, args...).Scan(&n); err != nil {
		t.Fatalf("%s: %v", query, err)
	}

	return n
}

func exists(t *testing.T, database *sql.DB, query string, args ...any) bool {
	t.Helper()

	return count(t, database, `SELECT count(*) FROM (`+query+`) q`, args...) > 0
}

func TestFreshDatabaseGetsAWorkingSchema(t *testing.T) {
	database := dbtest.NewMigrated(t)

	for _, table := range []string{"admin_config", "group", "country", "score"} {
		if !exists(t, database, `SELECT 1 FROM information_schema.tables WHERE table_schema = current_schema() AND table_name = $1`, table) {
			t.Errorf("table %s is missing", table)
		}
	}
	if exists(t, database, `SELECT 1 FROM information_schema.tables WHERE table_schema = current_schema() AND table_name = 'auth'`) {
		t.Error("the unused auth table must not be created")
	}

	for _, index := range []string{"score_user_country_year_key", "score_user_year_idx", "country_year_name_key", "group_owner_name_key"} {
		if !exists(t, database, `SELECT 1 FROM pg_indexes WHERE indexname = $1`, index) {
			t.Errorf("index %s is missing", index)
		}
	}
	if !exists(t, database, `SELECT 1 FROM information_schema.columns WHERE table_name = 'group' AND column_name = 'members' AND data_type = 'ARRAY'`) {
		t.Error("group.members must be a real array")
	}

	for year, want := range map[int]int{2024: 37, 2025: 37, 2026: 34} {
		if got := count(t, database, `SELECT count(*) FROM country WHERE year = $1`, year); got != want {
			t.Errorf("countries in %d = %d, want %d", year, got, want)
		}
	}
	if got := count(t, database, `SELECT count(*) FROM admin_config WHERE id = 1 AND year = 2026`); got != 1 {
		t.Errorf("admin_config seed rows = %d, want 1", got)
	}
	if _, err := database.Exec(`INSERT INTO admin_config (id, year, gameType, isVotingActive) VALUES (2, 2027, 'semi1', false)`); err == nil {
		t.Error("a second admin_config row must be rejected")
	}
}

func TestSecondRunChangesNothing(t *testing.T) {
	database := dbtest.NewMigrated(t)

	snapshot := func() [4]int {
		return [4]int{
			count(t, database, `SELECT count(*) FROM goose_db_version`),
			count(t, database, `SELECT count(*) FROM country`),
			count(t, database, `SELECT count(*) FROM admin_config`),
			count(t, database, `SELECT count(*) FROM pg_indexes WHERE schemaname = current_schema()`),
		}
	}
	before := snapshot()

	provider, err := migrations.NewProvider(database, true)
	if err != nil {
		t.Fatal(err)
	}
	results, err := provider.Up(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 0 {
		t.Fatalf("second run applied %d migrations, want 0", len(results))
	}
	if err := migrations.Run(context.Background(), database); err != nil {
		t.Fatal(err)
	}

	if after := snapshot(); after != before {
		t.Fatalf("second run changed the database: before %v, after %v", before, after)
	}
}

func TestConcurrentStartsApplyEachMigrationOnce(t *testing.T) {
	first, url := dbtest.NewDatabase(t)

	second, err := sql.Open("postgres", url)
	if err != nil {
		t.Fatal(err)
	}
	defer second.Close()

	var wg sync.WaitGroup
	errs := make(chan error, 2)
	for _, database := range []*sql.DB{first, second} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			errs <- migrations.Run(context.Background(), database)
		}()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			t.Fatalf("concurrent run failed: %v", err)
		}
	}

	if got := count(t, first, `SELECT count(*) FROM goose_db_version WHERE version_id > 0`); got != 6 {
		t.Fatalf("applied versions = %d, want 6 (each exactly once)", got)
	}
	if got := count(t, first, `SELECT count(*) FROM country`); got != 108 {
		t.Fatalf("countries = %d, want 108 (seeded once)", got)
	}
}

func TestExistingDatabaseIsAdoptedAndDeduplicated(t *testing.T) {
	database, _ := dbtest.NewDatabase(t)

	provider, err := migrations.NewProvider(database, true)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	if _, err := provider.UpTo(ctx, 1); err != nil {
		t.Fatal(err)
	}

	// what the old code could leave behind
	for _, stmt := range []string{
		`INSERT INTO score ("user", country, year, gametype, infinal, position) VALUES
			('bob', 'France', 2026, 'final', false, 0),
			('bob', 'France', 2026, 'final', true, 3),
			('bob', 'France', 2026, 'final', false, 0),
			('amy', 'France', 2026, 'final', false, 0)`,
		`INSERT INTO country (name, code, year, gametype, score, isinfinal, artist, song, ordersemi, orderfinal)
			VALUES ('France', 'fr', 2026, 'final', 7, true, 'Live artist', 'Live song', 0, 4),
			       ('France', 'fr', 2026, 'final', 0, false, 'Stale', 'Stale', 0, 0)`,
		`INSERT INTO "group" (name, owner, members, datecreated) VALUES
			('Friends', 'bob', '{bob}', now()), ('Friends', 'bob', '{bob,amy}', now())`,
	} {
		if _, err := database.Exec(stmt); err != nil {
			t.Fatalf("%s: %v", stmt, err)
		}
	}

	if err := migrations.Run(ctx, database); err != nil {
		t.Fatal(err)
	}

	if got := count(t, database, `SELECT count(*) FROM score WHERE "user" = 'bob'`); got != 1 {
		t.Errorf("bob's duplicate scores = %d rows, want 1", got)
	}
	if !exists(t, database, `SELECT 1 FROM score WHERE "user" = 'bob' AND infinal AND position = 3`) {
		t.Error("the row with the most progress must survive")
	}
	if got := count(t, database, `SELECT count(*) FROM score`); got != 2 {
		t.Errorf("score rows = %d, want 2", got)
	}

	// the live row survives and the seed does not overwrite it
	if got := count(t, database, `SELECT count(*) FROM country WHERE name = 'France' AND year = 2026`); got != 1 {
		t.Errorf("France 2026 rows = %d, want 1", got)
	}
	if !exists(t, database, `SELECT 1 FROM country WHERE name = 'France' AND year = 2026 AND score = 7 AND artist = 'Live artist'`) {
		t.Error("existing country data must not be overwritten by the seed")
	}

	if got := count(t, database, `SELECT count(DISTINCT name) FROM "group" WHERE owner = 'bob'`); got != 2 {
		t.Errorf("duplicate group names were not made unique")
	}
}

func TestOldTextMembersColumnIsRecreatedAsArray(t *testing.T) {
	database, _ := dbtest.NewDatabase(t)

	provider, err := migrations.NewProvider(database, true)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	if _, err := provider.UpTo(ctx, 1); err != nil {
		t.Fatal(err)
	}
	for _, stmt := range []string{
		`DROP TABLE "group"`,
		`CREATE TABLE "group" (id SERIAL PRIMARY KEY, name VARCHAR(20) NOT NULL, owner VARCHAR(50) NOT NULL, members TEXT NOT NULL, datecreated TIMESTAMP NOT NULL)`,
		`INSERT INTO "group" (name, owner, members, datecreated) VALUES ('Old', 'bob', '{bob}', now())`,
	} {
		if _, err := database.Exec(stmt); err != nil {
			t.Fatal(err)
		}
	}

	if err := migrations.Run(ctx, database); err != nil {
		t.Fatal(err)
	}

	if !exists(t, database, `SELECT 1 FROM information_schema.columns WHERE table_name = 'group' AND column_name = 'members' AND data_type = 'ARRAY'`) {
		t.Error("members must be a real array after migrating")
	}
	if got := count(t, database, `SELECT count(*) FROM "group"`); got != 0 {
		t.Errorf("groups reset with the recreated table, got %d rows", got)
	}
}
