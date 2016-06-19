package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (r *ClassReader) readUint8() uint8 {
	val := r.data[0]
	r.data = r.data[1:]
	return val
}

func (r *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(r.data)
	r.data = r.data[2:]
	return val
}

func (r *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(r.data)
	r.data = r.data[4:]
	return val
}

func (r *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(r.data)
	r.data = r.data[8:]
	return val
}

func (r *ClassReader) readUint16s() []uint16 {
	n := r.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = r.readUint16()
	}
	return s
}

func (r *ClassReader) readBytes(n uint32) []byte {
	bytes := r.data[:n]
	r.data = r.data[n:]
	return bytes
}
