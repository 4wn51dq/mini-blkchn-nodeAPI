package state

import (
	"net/http"

	"github.com/ethereum/go-ethereum/rlp"
)

// use go-ethereum rlp library

type (
	raw rlp.RawValue
)

func (raw *raw) Decode(w http.ResponseWriter, r *http.Request) error {
	return rlp.Decode(r.Body)
}
