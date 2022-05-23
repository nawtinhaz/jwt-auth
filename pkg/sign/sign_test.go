package sign_test

import (
	"testing"

	"github.com/nawtinhaz/jwt-auth/pkg/hash"
	"github.com/nawtinhaz/jwt-auth/pkg/sign"
)

func TestSign(t *testing.T) {
	msg := "Pasteld3Na1aehOMelh0rDoM()nd03M4is4lém"
	key, _ := hash.Keygen()
	sign.Sign([]byte(msg), key)
}

func TestValid(t *testing.T) {
	msg := "Pasteld3Na1aehOMelh0rDoM()nd03M4is4lém"
	key, _ := hash.Keygen()
	signature := sign.Sign([]byte(msg), key)

	tcs := []struct {
		name string
		msg  []byte
		sig  []byte
		exp  error
	}{
		{
			name: "message valid",
			msg:  []byte(msg),
			sig:  signature,
			exp:  nil,
		},
		{
			name: "message invalid",
			msg:  []byte("anothermessage"),
			sig:  signature,
			exp:  sign.ErrInvalidSignature,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if err := sign.Valid(tc.msg, tc.sig, key); err != tc.exp {
				t.Errorf("unexpected result when comparing macs: got %v, exp: %v", err, tc.exp)
			}
		})
	}
}
