package p2p

import (
	"encoding/binary"
	"io"
)

func WriteMessage(w io.Writer, t MessageType, payload []byte) error {

	if err := binary.Write(w, binary.BigEndian, t); err != nil {
		return err
	}

	length := uint32(len(payload))
	if err := binary.Write(w, binary.BigEndian, length); err != nil {
		return err
	}

	_, err := w.Write(payload)
	return err
}

func ReadMessage(r io.Reader) (MessageType, []byte, error) {

	var t MessageType
	if err := binary.Read(r, binary.BigEndian, &t); err != nil {
		return 0, nil, err
	}

	var length uint32
	if err := binary.Read(r, binary.BigEndian, &length); err != nil {
		return 0, nil, err
	}

	payload := make([]byte, length)
	_, err := io.ReadFull(r, payload)
	return t, payload, err
}