package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	register(controltype.PublishAck, NewPublishAck)
}

// PublishAck PublishAck
type PublishAck struct {
	flags uint8
}

// NewPublishAck NewPublishAck
func NewPublishAck(flags uint8) Message {
	return &PublishAck{
		flags: flags,
	}
}

// Decode Decode
func (m *PublishAck) Decode(br *bufio.Reader) error {
	return nil
}

// Encode Encode
func (m *PublishAck) Encode(bw *bufio.Writer) error {
	return nil
}
