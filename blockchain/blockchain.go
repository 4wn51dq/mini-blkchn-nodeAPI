package blockchain

import "time"

type Tx struct {
	From      string
	To        string
	Amount    string
	Timestamp string
}

type Block struct {
	Height       uint
	Timestamp    time.Time
	PrevHash     string
	MerkleRoot   string
	Transactions []Tx
}

type Blockchain struct {
	Chain []Block
}

func NewBlockchain() *Blockchain {
	genesis := Block{
		Height:     0,
		Timestamp:  time.Now(),
		MerkleRoot: "genesis",
	}
	return &Blockchain{
		Chain: []Block{genesis},
	}
}

func LastBlock(bc *Blockchain) Block {
	return bc.Chain[len(bc.Chain)-1]
}

func AddBlock(bc *Blockchain, txs []Tx) Blockchain {
	prev := bc.Chain[len(bc.Chain)-1]
	block := Block{
		Height:       prev.Height + 1,
		Timestamp:    time.Now(),
		PrevHash:     prev.PrevHash,
		MerkleRoot:   "someHash",
		Transactions: txs,
	}
	bc.Chain = append(bc.Chain, block)
	return *bc
}
