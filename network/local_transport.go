package network

import (
	"fmt"
	"sync"
)

//Local tranport is going to be used for testing, boot up nodes without managing tcp/..
type LocalTransport struct {
	address   NetAddress
	consumeCh chan RPC
	peers     map[NetAddress]*LocalTransport
	lock      sync.RWMutex
}

func NewLocalTransport(address NetAddress) *LocalTransport {
	return &LocalTransport{
		address:   address,
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[NetAddress]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeCh
}

func (t *LocalTransport) Connect(tr *LocalTransport) error {

	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr
	return nil
}

func (t *LocalTransport) SendMessage(to NetAddress, payload []byte) error {

	t.lock.RLock()
	defer t.lock.RUnlock()

	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf(" %s Could not send message to %s", t.Addr(), to)
	}

	peer.consumeCh <- RPC{
		From: t.Addr(),
		Payload: payload,
	}
	return nil
}


func (t *LocalTransport) Addr() NetAddress {
	return t.address
}