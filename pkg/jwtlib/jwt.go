package jwtlib

import (
	"encoding/json"
	"time"

	"github.com/cristalhq/jwt/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// NewToken creates a new jwt from the given key.
func NewToken(key []byte) (*jwt.Token, error) {
	signer := NewSignerHS(key)

	tokenID, err := gonanoid.Generate()
	if err != nil {
		return nil, err
	}
	claims := jwt.RegisteredClaims{
		Audience: []string{"nawtinhaz"},
		ID:       tokenID,
	}

	builder := jwt.NewBuilder(&signer)
	token, err := builder.Build(claims)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// VerifyClaims verifies that the token claims are valid.
func VerifyClaims(key []byte, tokenBytes []byte) error {
	verifier := NewVerifierHS(key)

	token, err := jwt.Parse(tokenBytes, verifier)
	if err != nil {
		return err
	}

	var claims jwt.RegisteredClaims
	if err := json.Unmarshal(token.Claims(), &claims); err != nil {
		return err
	}

	// or parse only claims
	if err := jwt.ParseClaims(tokenBytes, verifier, &claims); err != nil {
		return err
	}

	if valid := claims.IsForAudience("nawtinhaz"); !valid {
		return ErrInvalidToken
	}
	if valid := claims.IsValidAt(time.Now()); !valid {
		return ErrInvalidToken
	}
	return nil
}
