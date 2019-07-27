package packet

import (
	"encoding/binary"
	"fmt"
)

var networkOrder = binary.BigEndian

type Writer struct {
	buffer []byte
	offset int
}

func NewWriter(buffer []byte) *Writer {
	return &Writer{buffer, 0}
}

func NewWriterSize(n int) *Writer {
	return NewWriter(make([]byte, n))
}

func (w *Writer) WriteByte(v byte) {
	w.buffer[w.offset] = v
	w.offset++
}

func (w *Writer) WriteUint16(v uint16) {
	networkOrder.PutUint16(w.buffer[w.offset:], v)
	w.offset += 2
}

func (w *Writer) WriteUint24(v uint32) {
	w.WriteByte(byte(v >> 16 & 0xff))
	w.WriteByte(byte(v >> 8 & 0xff))
	w.WriteByte(byte(v & 0xff))
}

func (w *Writer) WriteUint32(v uint32) {
	networkOrder.PutUint32(w.buffer[w.offset:], v)
	w.offset += 4
}

// Write the given bytes, if there is enough room.
func (w *Writer) WriteSlice(p []byte) error {
	if err := w.CheckCapacity(len(p)); err != nil {
		return err
	}
	w.offset += copy(w.buffer[w.offset:], p)
	return nil
}

// Return the number of bytes written so far.
func (w *Writer) Length() int {
	return w.offset
}

func (w *Writer) Rewind(n int) {
	w.offset -= n
	if w.offset < 0 {
		w.offset = 0
	}
}

// Return the number of bytes that the underlying buffer can hold.
func (w *Writer) Capacity() int {
	return len(w.buffer)
}

func (w *Writer) CheckCapacity(needed int) error {
	if w.Capacity() < needed {
		return fmt.Errorf("%d bytes available, %d needed", w.Capacity(), needed)
	}
	return nil
}

// Return a slice of the bytes written so far.
func (w *Writer) Bytes() []byte {
	return w.buffer[0:w.offset]
}

func (w *Writer) Reset() {
	w.offset = 0
}
