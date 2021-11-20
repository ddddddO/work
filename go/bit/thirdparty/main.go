package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/nokute78/go-bit/v2"
)

func main() {
	type TcpHeader struct {
		SrcPort    uint16
		DstPort    uint16
		SeqNo      uint32
		AckNo      uint32
		HeaderLen  [4]bit.Bit
		Reserved   [3]bit.Bit `bit:"skip"`
		NS         bit.Bit
		CWR        bit.Bit
		ECE        bit.Bit
		URG        bit.Bit
		ACK        bit.Bit
		PSH        bit.Bit
		RST        bit.Bit
		SYN        bit.Bit
		FIN        bit.Bit
		WinSize    uint16
		CheckSum   uint16
		EmePointer uint16
	}

	s := TcpHeader{}

	br := bytes.NewReader([]byte{0xd8, 0x65, 0x01, 0xbb, 0x4b, 0xe0, 0x76, 0xcd, 0x48, 0xc8, 0x70, 0x8f, 0x50, 0x10, 0x10,
		0x18, 0x0e, 0xc1, 0x00, 0x00})

	if err := bit.Read(br, binary.BigEndian, &s); err != nil {
		fmt.Printf("error:%s", err)
	}

	fmt.Printf("src=%d dst=%d\n", s.SrcPort, s.DstPort)
	fmt.Printf("SeqNo=%d AckNo=%d\n", s.SeqNo, s.AckNo)
	fmt.Printf("HeaderLen(raw)=%v\n", s.HeaderLen)
	fmt.Printf("Ack=%t\n", s.ACK)
}
