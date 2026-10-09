package group

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/vilmis04/eurovision-game-service/internal/auth"
)

type fakeStore struct {
	groups    map[int64]*Group
	err       error
	createErr error
	updates   map[int64][]string
	deleted   []int64
}

func newFakeStore() *fakeStore {
	return &fakeStore{
		groups: map[int64]*Group{
			1: {Id: 1, Name: "Friends", Owner: "alice", Members: []string{"alice", "bob"}},
		},
		updates: map[int64][]string{},
	}
}

func (f *fakeStore) GetGroupList(user string, groupId string) (*[]Group, error) {
	if f.err != nil {
		return nil, f.err
	}
	return &[]Group{}, nil
}

func (f *fakeStore) GetGroupById(id int64) (*Group, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.groups[id], nil
}

func (f *fakeStore) GetGroupNames(owner string) (*[]string, error) {
	return &[]string{"Friends"}, nil
}

func (f *fakeStore) CreateGroup(group *Group) (*int64, error) {
	if f.createErr != nil {
		return nil, f.createErr
	}
	id := int64(2)
	return &id, nil
}

func (f *fakeStore) UpdateMembers(id int64, members []string) error {
	f.updates[id] = members
	return nil
}

func (f *fakeStore) DeleteGroup(owner string, id int64) error {
	f.deleted = append(f.deleted, id)
	return nil
}

func newTestApp(store *fakeStore) *gin.Engine {
	gin.SetMode(gin.TestMode)
	app := gin.New()
	app.Use(auth.Proxy("secret"))
	ctrl := &controller{
		service: &Service{
			store:        store,
			inviteSecret: testSecret,
			now:          func() time.Time { return testNow },
		},
		router: app.Group("api/group"),
	}
	ctrl.Use()

	return app
}

func call(app *gin.Engine, method string, path string, user string, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set(auth.HeaderInternalToken, "secret")
	req.Header.Set(auth.HeaderUser, user)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	return rec
}

func TestPatchRequiresGroupOwner(t *testing.T) {
	store := newFakeStore()
	app := newTestApp(store)

	rec := call(app, http.MethodPatch, "/api/group/1", "bob", `{"members":["mallory"]}`)
	if rec.Code != http.StatusForbidden {
		t.Fatalf("member (not owner) status = %d, want 403", rec.Code)
	}
	rec = call(app, http.MethodPatch, "/api/group/1", "mallory", `{"members":["mallory"]}`)
	if rec.Code != http.StatusForbidden {
		t.Fatalf("stranger status = %d, want 403", rec.Code)
	}
	if len(store.updates) != 0 {
		t.Fatalf("forbidden requests must not update anything, got %v", store.updates)
	}
}

func TestPatchRejectsMalformedRequests(t *testing.T) {
	app := newTestApp(newFakeStore())

	tests := []struct {
		name, path, body string
		status           int
	}{
		{"non numeric id", "/api/group/abc", `{"members":["x"]}`, http.StatusBadRequest},
		{"sql in id", "/api/group/1%20OR%201=1", `{"members":["x"]}`, http.StatusBadRequest},
		{"negative id", "/api/group/-1", `{"members":["x"]}`, http.StatusBadRequest},
		{"unknown group", "/api/group/99", `{"members":["x"]}`, http.StatusNotFound},
		{"broken json", "/api/group/1", `{`, http.StatusBadRequest},
		{"blank member", "/api/group/1", `{"members":["  "]}`, http.StatusBadRequest},
		{"huge member", "/api/group/1", `{"members":["` + strings.Repeat("a", 300) + `"]}`, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := call(app, http.MethodPatch, tt.path, "alice", tt.body)
			if rec.Code != tt.status {
				t.Fatalf("status = %d, want %d (body %q)", rec.Code, tt.status, rec.Body.String())
			}
		})
	}
}

func TestPatchByOwnerAddsMembersOnce(t *testing.T) {
	store := newFakeStore()
	app := newTestApp(store)

	rec := call(app, http.MethodPatch, "/api/group/1", "alice", `{"members":["carol","bob","carol"]}`)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if want := []string{"alice", "bob", "carol"}; !slices.Equal(store.updates[1], want) {
		t.Fatalf("members = %v, want %v", store.updates[1], want)
	}
}

func TestGenerateInviteRequiresGroupOwner(t *testing.T) {
	app := newTestApp(newFakeStore())

	rec := call(app, http.MethodPost, "/api/group/1/generate-invite", "bob", "")
	if rec.Code != http.StatusForbidden {
		t.Fatalf("member status = %d, want 403", rec.Code)
	}
	rec = call(app, http.MethodPost, "/api/group/99/generate-invite", "alice", "")
	if rec.Code != http.StatusNotFound {
		t.Fatalf("unknown group status = %d, want 404 (used to panic)", rec.Code)
	}
	rec = call(app, http.MethodPost, "/api/group/abc/generate-invite", "alice", "")
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("bad id status = %d, want 400", rec.Code)
	}
	rec = call(app, http.MethodPost, "/api/group/1/generate-invite", "alice", "")
	if rec.Code != http.StatusCreated {
		t.Fatalf("owner status = %d, want 201", rec.Code)
	}
	if id, err := parseInvite(testSecret, rec.Body.String(), testNow); err != nil || id != 1 {
		t.Fatalf("invite id = %d, err = %v", id, err)
	}
}

func TestJoinRejectsInvalidInvites(t *testing.T) {
	store := newFakeStore()
	app := newTestApp(store)

	expired, _ := createInvite(testSecret, 1, testNow.Add(-time.Second))
	forged, _ := createInvite([]byte("guess"), 1, testNow.Add(InviteTTL))

	tests := []struct{ name, body string }{
		{"broken json", `{`},
		{"missing code", `{}`},
		{"garbage", `{"inviteCode":"abc"}`},
		{"legacy base64 invite", `{"inviteCode":"RnJpZW5kczphbGljZToxOjIwMjQtMDEtMDE"}`},
		{"expired", `{"inviteCode":"` + expired + `"}`},
		{"forged signature", `{"inviteCode":"` + forged + `"}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := call(app, http.MethodPost, "/api/group/join", "mallory", tt.body)
			if rec.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want 400 (body %q)", rec.Code, rec.Body.String())
			}
		})
	}
	if len(store.updates) != 0 {
		t.Fatalf("rejected invites must not change membership, got %v", store.updates)
	}
}

func TestJoinWithValidInvite(t *testing.T) {
	store := newFakeStore()
	app := newTestApp(store)
	invite, _ := createInvite(testSecret, 1, testNow.Add(InviteTTL))
	body := `{"inviteCode":"` + invite + `"}`

	rec := call(app, http.MethodPost, "/api/group/join", "carol", body)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if want := []string{"alice", "bob", "carol"}; !slices.Equal(store.updates[1], want) {
		t.Fatalf("members = %v, want %v", store.updates[1], want)
	}

	delete(store.updates, 1)
	rec = call(app, http.MethodPost, "/api/group/join", "bob", body)
	if rec.Code != http.StatusOK || len(store.updates) != 0 {
		t.Fatalf("existing member: status = %d, updates = %v", rec.Code, store.updates)
	}

	store.groups = map[int64]*Group{}
	rec = call(app, http.MethodPost, "/api/group/join", "carol", body)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("deleted group status = %d, want 404 (used to panic)", rec.Code)
	}
}

func TestInternalErrorsAreNotExposed(t *testing.T) {
	store := newFakeStore()
	store.err = errors.New(`pq: password authentication failed for user "secret_db_user"`)
	app := newTestApp(store)

	rec := call(app, http.MethodGet, "/api/group/", "alice", "")
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", rec.Code)
	}
	if strings.Contains(rec.Body.String(), "secret_db_user") || strings.Contains(rec.Body.String(), "pq:") {
		t.Fatalf("internal error leaked: %q", rec.Body.String())
	}
}

func TestCreateGroupRejectsInvalidNames(t *testing.T) {
	app := newTestApp(newFakeStore())

	tests := []struct {
		name, body string
		status     int
	}{
		{"broken json", `{`, http.StatusBadRequest},
		{"empty", `{"name":" "}`, http.StatusBadRequest},
		{"too long", `{"name":"` + strings.Repeat("a", 21) + `"}`, http.StatusBadRequest},
		{"duplicate", `{"name":"Friends"}`, http.StatusConflict},
		{"valid", `{"name":"Fresh"}`, http.StatusCreated},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := call(app, http.MethodPost, "/api/group/", "alice", tt.body)
			if rec.Code != tt.status {
				t.Fatalf("status = %d, want %d", rec.Code, tt.status)
			}
		})
	}
}

func TestGroupEndpointsRejectSpoofedUser(t *testing.T) {
	app := newTestApp(newFakeStore())

	req := httptest.NewRequest(http.MethodPatch, "/api/group/1", strings.NewReader(`{"members":["x"]}`))
	req.Header.Set(auth.HeaderUser, "alice") // no proxy token
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", rec.Code)
	}
}

func TestCreateGroupMapsUniqueViolationToConflict(t *testing.T) {
	store := newFakeStore()
	store.createErr = &pq.Error{Code: "23505", Message: "duplicate key value violates unique constraint group_owner_name_key"}
	app := newTestApp(store)

	rec := call(app, http.MethodPost, "/api/group/", "alice", `{"name":"Fresh"}`)
	if rec.Code != http.StatusConflict {
		t.Fatalf("status = %d, want 409", rec.Code)
	}
	if strings.Contains(rec.Body.String(), "group_owner_name_key") {
		t.Fatalf("constraint name leaked: %q", rec.Body.String())
	}
}
