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
		exp  bool
	}{
		{
			name: "message valid",
			msg:  []byte(msg),
			sig:  signature,
			exp:  true,
		},
		{
			name: "message invalid",
			msg:  []byte("anothermessage"),
			sig:  signature,
			exp:  false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if got, exp := sign.Valid(tc.msg, tc.sig, key), tc.exp; got != exp {
				t.Errorf("unexpected value when comparing macs: got %v, exp: %v", got, exp)
			}
		})
	}
}
