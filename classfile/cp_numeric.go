package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

func (c *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	c.val = int32(bytes)
}

func (self *ConstantIntegerInfo) Value() int32 {
	return self.val
}

type ConstantFloatInfo struct {
	val float32
}

func (c *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	c.val = math.Float32frombits(bytes)
}

func (self *ConstantFloatInfo) Value() float32 {
	return self.val
}

type ConstantLongInfo struct {
	val int64
}

func (c *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	c.val = int64(bytes)
}

func (self *ConstantLongInfo) Value() int64 {
	return self.val
}

type ConstantDoubleInfo struct {
	val float64
}

func (c *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	c.val = math.Float64frombits(bytes)
}

func (self *ConstantDoubleInfo) Value() float64 {
	return self.val
}
