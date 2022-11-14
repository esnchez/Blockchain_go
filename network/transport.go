package network

type NetAddress string

//message that is going to be sent over the transport layer
type RPC struct {
	From    NetAddress
	Payload []byte
}

type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(NetAddress, []byte) error
	Addr() NetAddress
}
