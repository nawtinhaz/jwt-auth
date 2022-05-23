package jwtlib

// ErrInvalidToken is returned when a token is not valid.
const ErrInvalidToken TokenErr = "the token payload is not valid"

type TokenErr string

func (e TokenErr) Error() string {
	return string(e)
}
