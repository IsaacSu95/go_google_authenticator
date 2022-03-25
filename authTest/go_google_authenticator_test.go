package authTest

import (
	"crypto/rand"
	"encoding/base32"
	"isaacSu/go_google_authenticator"
	"testing"
)

var s = "CA5Z4EJD4OBB7YSK"

func TestGetGoogleAuth(t *testing.T) {
	t.Log(go_google_authenticator.GetGoogleAuth(s))
}

func TestGoogleAuth(t *testing.T) {
	t.Log(go_google_authenticator.GoogleAuth(s, "719366"))
}

func TestNewGoogleAuth(t *testing.T) {
	key := go_google_authenticator.RandNewGoogleAuthKey()
	t.Log(key, len(key))
}

func TestBase32(t *testing.T) {
	d := make([]byte, 10)
	t.Log(rand.Read(d))
	t.Log(string(d))
	b := base32.StdEncoding
	t.Log(b.EncodeToString(d))
}
