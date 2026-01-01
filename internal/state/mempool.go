package state

import "sync"

type Mempool struct {
	mu  sync.Mutex
	txs []*Tx
}

func (m *Mempool) AddToMempool(tx *Tx) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.txs = append(m.txs, tx)
}
