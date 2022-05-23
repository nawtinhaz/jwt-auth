package hash

// ErrMismatchedHashAndPassword is returned when a password (hashed) and
// given hash do not match.
const ErrMismatchedHashandPassword MismatchErr = "the hashed password does not match the hash of the given password"

type MismatchErr string

func (e MismatchErr) Error() string {
	return string(e)
}
