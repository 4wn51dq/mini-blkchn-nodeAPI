package state

import (
	"encoding/json"
	"io"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	LEGACY_TYPE byte = 0x1
	EIP2930     byte = 0x2
	EIP1559     byte = 0x3
)

type Block struct {
	header Header
	txs    Txs
}

type Header struct {
	difficulty      *big.Int
	totalDifficulty *big.Int
	extraData       []byte
	gasLimit        uint
	gasUsed         uint
	hash            common.Hash
	logsBloom       []byte
	miner           string
	mixhash         []byte
	nonce           uint
	number          uint
	size            uint
	timestamp       time.Time
	sha3uncles      []byte
	uncles          [][]byte
	parentHash      []byte
	txsRoot         []byte
	receiptRoot     []byte
	stateRoot       []byte
}

type Txs []*Tx

type Signature struct {
	V, R, S *big.Int
}

type AccessTuple struct {
	Address     common.Address
	StorageKeys [][32]byte
}

type AccessList []AccessTuple

type Tx struct {
	chainId              *big.Int
	txType               byte
	nonce                uint64
	recipient            common.Address // 20 bytes
	value                *big.Int
	yParity, r, s        *big.Int
	calldata             []byte
	gasLimit             uint64
	baseFeePerGas        *big.Int
	maxPriorityFeePerGas *big.Int
	maxFeePerGas         *big.Int
	accessList           AccessList
}

// RecordTx collects tx in mempool
func (m Mempool) RecordTx(tx *Tx) {
	m.AddToMempool(tx)
}

func (tx *Tx) TxFromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(tx)
}
