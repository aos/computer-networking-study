package main

import (
	"bytes"
	"encoding/binary"
)

type icmp struct {
	Type       uint8
	Code       uint8
	Checksum   uint16
	Identifier uint16
	Sequence   uint16
	Timestamp  int64
}

// AssembleICMP takes an ICMP type, code, identifier, sequence number and
// timestamp and returns an assembled byte buffer
func AssembleICMP(t, c uint8, id, seq uint16, ts int64) []byte {
	icmpPacket := icmp{
		Type:       t,
		Code:       c,
		Checksum:   0,
		Identifier: id,
		Sequence:   seq,
		Timestamp:  ts,
	}

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmpPacket)

	icmpPacket.Checksum = Checksum(buffer.Bytes())
	buffer.Reset()
	binary.Write(&buffer, binary.BigEndian, icmpPacket)
	return buffer.Bytes()
}
