package bytesutil

import (
	"encoding/binary"
	"sync"
)

var pool = &sync.Pool{
	New: func() interface{} {
		return New()
	},
}

// Get Get
func Get(data []byte) *Bytes {
	bytes := pool.Get().(*Bytes)
	bytes.Reset(data)
	return bytes
}

// Put Put
func Put(bytes *Bytes) {
	pool.Put(bytes)
}

// Bytes Bytes
type Bytes struct {
	data   []byte
	cursor int64
}

// New New
func New() *Bytes {
	return &Bytes{}
}

// ReadUint16 ReadUint16
func (b *Bytes) ReadUint16() uint16 {
	result := binary.BigEndian.Uint16(b.data[b.cursor : b.cursor+2])
	b.cursor += 2
	return result
}

// ReadUint32 ReadUint32
func (b *Bytes) ReadUint32() uint32 {
	result := binary.BigEndian.Uint32(b.data[b.cursor : b.cursor+4])
	b.cursor += 4
	return result
}

// ReadByte ReadByte
func (b *Bytes) ReadByte() byte {
	return b.Read(1)[0]
}

// Read Read
func (b *Bytes) Read(dataLen uint16) []byte {
	data := b.data[b.cursor : b.cursor+int64(dataLen)]
	b.cursor += int64(dataLen)
	return data
}

// ReadVariableStr ReadVariableStr
func (b *Bytes) ReadVariableStr() string {
	return string(b.ReadVariable())
}

// ReadVariable ReadVariable
func (b *Bytes) ReadVariable() []byte {
	dataLen := b.ReadUint16()
	return b.Read(dataLen)
}

// ReadStr ReadStr
func (b *Bytes) ReadStr(dataLen uint16) string {
	return string(b.Read(dataLen))
}

// ReadAll ReadAll
func (b *Bytes) ReadAll() []byte {
	data := b.data[b.cursor:]
	b.cursor = int64(len(b.data))
	return data
}

// ReadAllStr ReadAllStr
func (b *Bytes) ReadAllStr() string {
	return string(b.ReadAll())
}

// EOF EOF
func (b *Bytes) EOF() bool {
	return b.cursor == int64(len(b.data))
}

// Reset Reset
func (b *Bytes) Reset(data []byte) {
	b.data = data
	b.cursor = 0
}
