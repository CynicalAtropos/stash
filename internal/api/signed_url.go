package api

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/stashapp/stash/internal/manager/config"
	"github.com/stashapp/stash/pkg/session"
	"github.com/stashapp/stash/pkg/signedurl"
)

// userSigningKey returns the HMAC signing key for a given user.
func userSigningKey(c *config.Config, _ string) []byte {
	return c.GetJWTSignKey()
}

// signedParams generates signed URL query parameters for the given path prefix and user.
func signedParams(c *config.Config, userID string, prefix string) url.Values {
	secret := userSigningKey(c, userID)
	cid := signedurl.GenerateCredentialID(secret, userID)
	expires := time.Now().Add(c.GetSignedURLExpiry())
	return signedurl.SignPrefix(prefix, secret, cid, expires)
}

func getSignedURLUserID(ctx context.Context, c *config.Config) (*string, error) {
	if !c.HasCredentials() {
		return nil, nil
	}

	userID := session.GetCurrentUserID(ctx)
	if userID == nil {
		return nil, fmt.Errorf("user ID not found")
	}

	return userID, nil
}

func setStreamURLAuthParams(c *config.Config, userID *string, u *url.URL) {
	if c.HasCredentials() {
		u.RawQuery = signedParams(c, *userID, signedurl.DerivePrefix(u.Path)).Encode()
		return
	}

	apiKey := c.GetAPIKey()
	if apiKey != "" {
		v := u.Query()
		v.Set("apikey", apiKey)
		u.RawQuery = v.Encode()
	}
}

// resolveCredentialID maps a credential ID back to a username and their signing key.
func resolveCredentialID(c *config.Config, cid string) (string, []byte, bool) {
	username := c.GetUsername()
	secret := userSigningKey(c, username)
	if signedurl.GenerateCredentialID(secret, username) == cid {
		return username, secret, true
	}
	return "", nil, false
}
