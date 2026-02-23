package p2p

import (
	"net"
	"sync"
)

type Server struct {
	ListenAddr string
	Peers      map[string]*Peer
	mu         sync.Mutex
}

func (s *Server) Start() error {

	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go s.handleConnection(conn)
	}
}