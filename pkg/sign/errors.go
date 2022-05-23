package sign

// ErrInvalidSignature is returned when a signature is not valid.
const ErrInvalidSignature SignatureErr = "the message signature is not valid"

type SignatureErr string

func (e SignatureErr) Error() string {
	return string(e)
}
