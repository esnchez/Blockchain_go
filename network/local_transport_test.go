package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestConnect(t *testing.T) {

	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)
	assert.Equal(t, tra.peers[trb.address], trb)
	assert.Equal(t, trb.peers[tra.address], tra)
}

func TestSendMessage(t *testing.T) {

	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)
	msg := []byte("hi")
	tra.SendMessage(trb.address,msg)

	rpc := <-trb.Consume()
	assert.Equal(t, rpc.From, tra.Addr())
	assert.Equal(t, rpc.Payload, msg)
}