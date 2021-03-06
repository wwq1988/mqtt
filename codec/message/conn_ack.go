package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	Register(controltype.ConnAck, NewConnAck)
}

// ConnAck ConnAck
type ConnAck struct {
	flags uint8
}

// NewConnAck NewConnAck
func NewConnAck(flags uint8) Message {
	return &ConnAck{
		flags: flags,
	}
}

// Decode Decode
func (m *ConnAck) Decode(data []byte) error {
	return nil
}

// EncodeTo EncodeTo
func (m *ConnAck) EncodeTo(br *bufio.Writer) error {
	return nil
}
