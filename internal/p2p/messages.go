package p2p

type MessageType uint8

const (
	MsgVersion MessageType = 1
	MsgVerAck  MessageType = 2
	MsgInv     MessageType = 3
	MsgGetData MessageType = 4
	MsgBlock   MessageType = 5
	MsgTx      MessageType = 6
)