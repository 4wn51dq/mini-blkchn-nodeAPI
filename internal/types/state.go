package state

import (
	"encoding/json"
	"net/http"
	"time"
)

type Block struct {
	header Header
	txs    Txs
}

type Header struct {
	difficulty      int
	totalDifficulty int
	extraData       string
	gasLimit        uint
	gasUsed         uint
	hash            []byte
	logsBloom       string
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

type Tx struct {
	chainId              int
	txType               byte
	nonce                uint
	recipient            []byte
	sender               []byte
	value                uint
	yParity, r, s        []byte
	calldata             string
	gasLimit             uint
	baseFeePerGas        *uint
	maxPriorityFeePerGas *uint
	maxFeePerGas         *uint
	accessList           [][]byte
}

// RecordTx collects tx in mempool
func RecordTx(tx *Tx) {
	tx.chainId = ETHEREUM_MAINNET_CHAINID
	tx.nonce = mempool.getNextNonce()
	mempool = append(mempool, tx)
}

func (tx Tx) TxFromJSON(w http.ResponseWriter, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(tx)
	return err
}

func (txs Txs) getNextNonce() uint {
	lastNonce := mempool[len(mempool)-1].nonce
	return lastNonce + 1
}

const (
	ETHEREUM_MAINNET_CHAINID int  = 1
	LEGACY_TYPE              byte = 0
	EIP2930                  byte = 2
	EIP1559                  byte = 4
)

var mempool Txs = []*Tx{
	&Tx{
		chainId:  1,
		txType:   0,
		nonce:    1,
		calldata: "genesis",
	},
}
