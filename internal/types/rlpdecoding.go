package state

import (
	"io"

	"github.com/ethereum/go-ethereum/rlp"
)

func (tx *Tx) getTxTypeRLPdecode(dst io.Writer) {
	encoder := rlp.NewEncoderBuffer(dst)
	encoder.AppendToBytes(tx.recipient)
}
