package claims

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrTokenExpired = errors.New("error validating token expiry attribute")
)

type Claims struct {
	jwt.RegisteredClaims
	SessionID int64 `json:"sessionID"`
}

func (c *Claims) Valid() error {
	if exp := c.VerifyExpiresAt(time.Now(), true); exp {
		return ErrTokenExpired
	}
	return nil
}
