package p2p

import (
	"net"
)

type Peer struct {
	Conn net.Conn
	Addr string
}