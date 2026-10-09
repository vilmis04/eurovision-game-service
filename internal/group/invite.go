package group

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// InviteTTL is how long a generated invite stays valid.
const InviteTTL = 7 * 24 * time.Hour

var (
	errInviteMalformed = errors.New("invite is malformed")
	errInviteSignature = errors.New("invite signature mismatch")
	errInviteExpired   = errors.New("invite expired")
	errInviteNoSecret  = errors.New("invite secret is not configured")
)

var inviteEncoding = base64.RawURLEncoding

func sign(secret []byte, payload []byte) []byte {
	mac := hmac.New(sha256.New, secret)
	mac.Write(payload)

	return mac.Sum(nil)
}

// createInvite returns `base64url(groupId:expiryUnix).base64url(hmac-sha256)`.
func createInvite(secret []byte, groupId int64, expiresAt time.Time) (string, error) {
	if len(secret) == 0 {
		return "", errInviteNoSecret
	}

	payload := []byte(fmt.Sprintf("%d:%d", groupId, expiresAt.Unix()))

	return inviteEncoding.EncodeToString(payload) + "." + inviteEncoding.EncodeToString(sign(secret, payload)), nil
}

// parseInvite checks the signature first and the expiry second, and returns the group id.
func parseInvite(secret []byte, invite string, now time.Time) (int64, error) {
	if len(secret) == 0 {
		return 0, errInviteNoSecret
	}

	parts := strings.Split(invite, ".")
	if len(parts) != 2 {
		return 0, errInviteMalformed
	}
	payload, err := inviteEncoding.DecodeString(parts[0])
	if err != nil {
		return 0, errInviteMalformed
	}
	signature, err := inviteEncoding.DecodeString(parts[1])
	if err != nil {
		return 0, errInviteMalformed
	}

	if !hmac.Equal(signature, sign(secret, payload)) {
		return 0, errInviteSignature
	}

	fields := strings.Split(string(payload), ":")
	if len(fields) != 2 {
		return 0, errInviteMalformed
	}
	groupId, err := strconv.ParseInt(fields[0], 10, 64)
	if err != nil {
		return 0, errInviteMalformed
	}
	expiry, err := strconv.ParseInt(fields[1], 10, 64)
	if err != nil {
		return 0, errInviteMalformed
	}
	if now.Unix() > expiry {
		return 0, errInviteExpired
	}

	return groupId, nil
}
