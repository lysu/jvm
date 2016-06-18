package base

type BytecodeReader struct {
	code []byte
	pc   int
}

func (self *BytecodeReader) PC() int {
	return self.pc
}

func (r *BytecodeReader) Reset(code []byte, pc int) {
	r.code = code
	r.pc = pc
}

func (r *BytecodeReader) ReadInt16() int16 {
	return int16(r.ReadUint16())
}

func (r *BytecodeReader) ReadUint16() uint16 {
	b1 := uint16(r.ReadUint8())
	b2 := uint16(r.ReadUint8())
	return (b1 << 8) | b2
}

func (r *BytecodeReader) ReadUint8() uint8 {
	i := r.code[r.pc]
	r.pc++
	return i
}

func (r *BytecodeReader) ReadInt8() int8 {
	return int8(r.ReadUint8())
}

func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}

func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUint8()
	}
}
