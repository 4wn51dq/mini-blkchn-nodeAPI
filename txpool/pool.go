package txpool

// in the mempool, tx wait until mined

type Tx struct {
	From      string
	To        string
	Amount    string
	Timestamp string
}

// var pool []Tx // why should i store this as a struct ?
type Pool struct {
	Transactions []Tx
}

func (p *Pool) Add(tx Tx) {
	p.Transactions = append(p.Transactions, tx)
}

func NewPool() *Pool {
	return &Pool{}
}
