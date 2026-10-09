package group

import (
	"strings"
	"testing"
	"time"
)

var (
	testSecret = []byte("invite-secret")
	testNow    = time.Date(2026, 5, 1, 12, 0, 0, 0, time.UTC)
)

func TestInviteRoundTrip(t *testing.T) {
	invite, err := createInvite(testSecret, 42, testNow.Add(InviteTTL))
	if err != nil {
		t.Fatal(err)
	}

	id, err := parseInvite(testSecret, invite, testNow.Add(time.Hour))
	if err != nil || id != 42 {
		t.Fatalf("id = %d, err = %v", id, err)
	}
}

func TestParseInviteRejects(t *testing.T) {
	valid, _ := createInvite(testSecret, 42, testNow.Add(InviteTTL))
	parts := strings.Split(valid, ".")
	otherSecret, _ := createInvite([]byte("other"), 42, testNow.Add(InviteTTL))
	forgedPayload, _ := createInvite(testSecret, 99, testNow.Add(InviteTTL))

	tests := []struct {
		name   string
		secret []byte
		invite string
		now    time.Time
		want   error
	}{
		{"empty", testSecret, "", testNow, errInviteMalformed},
		{"no signature", testSecret, parts[0], testNow, errInviteMalformed},
		{"too many parts", testSecret, valid + ".x", testNow, errInviteMalformed},
		{"not base64", testSecret, "!!!.???", testNow, errInviteMalformed},
		{"legacy unsigned invite", testSecret, "TmFtZTpvd25lcjoxOjIwMjQtMDEtMDE", testNow, errInviteMalformed},
		{"signed with another secret", testSecret, otherSecret, testNow, errInviteSignature},
		{"payload swapped under old signature", testSecret, strings.Split(forgedPayload, ".")[0] + "." + parts[1], testNow, errInviteSignature},
		{"tampered signature", testSecret, parts[0] + "." + parts[1][:len(parts[1])-2] + "AA", testNow, errInviteSignature},
		{"expired", testSecret, valid, testNow.Add(InviteTTL + time.Second), errInviteExpired},
		{"no secret configured", nil, valid, testNow, errInviteNoSecret},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parseInvite(tt.secret, tt.invite, tt.now)
			if err != tt.want {
				t.Fatalf("err = %v, want %v", err, tt.want)
			}
		})
	}
}

func TestCreateInviteRequiresSecret(t *testing.T) {
	if _, err := createInvite(nil, 1, testNow); err != errInviteNoSecret {
		t.Fatalf("err = %v, want errInviteNoSecret", err)
	}
}
