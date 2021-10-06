package sign_test

import (
	"testing"

	"github.com/nawtinhaz/jwt-auth/pkg/hash"
	"github.com/nawtinhaz/jwt-auth/pkg/sign"
)

func TestSign(t *testing.T) {
	message := "Pasteld3Na1aehOMelh0rDoM()nd03M4is4lém"
	key, _ := hash.Keygen()
	sign.Sign([]byte(message), key)
}

func TestValid(t *testing.T) {
	message := "Pasteld3Na1aehOMelh0rDoM()nd03M4is4lém"
	key, _ := hash.Keygen()
	messageSigned := sign.Sign([]byte(message), key)

	tcs := []struct {
		name          string
		message       []byte
		messageSigned []byte
		exp           bool
	}{
		{
			name:          "message valid",
			message:       []byte(message),
			messageSigned: messageSigned,
			exp:           true,
		},
		{
			name:          "message invalid",
			message:       []byte("anothermessage"),
			messageSigned: messageSigned,
			exp:           false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if got, exp := sign.Valid(tc.message, tc.messageSigned, key), tc.exp; got != exp {
				t.Errorf("unexpected value when comparing macs: got %v, exp: %v", got, exp)
			}
		})
	}
}
