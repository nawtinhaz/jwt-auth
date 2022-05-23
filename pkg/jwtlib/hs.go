package jwtlib

import (
	"crypto"

	"github.com/cristalhq/jwt/v4"
	"github.com/nawtinhaz/jwt-auth/pkg/sign"
)

type HSAlg struct {
	alg  jwt.Algorithm
	hash crypto.Hash
	key  []byte
}

func NewSignerHS(key []byte) HSAlg {
	return newHS(key)
}

func NewVerifierHS(key []byte) HSAlg {
	return newHS(key)
}

func (hs *HSAlg) Algorithm() jwtlib.Algorithm {
	return hs.alg
}

func (hs *HSAlg) SignSize() int {
	return hs.hash.Size()
}

func (hs *HSAlg) Sign(payload []byte) ([]byte, error) {
	return sign.Sign(payload, hs.key), nil
}

func (hs *HSAlg) Verify(token jwtlib.Token) error {
	return sign.Valid(token.PayloadPart(), token.SignaturePart(), hs.key)
}

func newHS(key []byte) HSAlg {
	return HSAlg{
		alg:  jwtlib.Algorithm("HS256"),
		hash: sign.SHA256,
		key:  key,
	}
}
