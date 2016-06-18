package classfile

type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (c *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	c.referenceKind = reader.readUint8()
	c.referenceIndex = reader.readUint16()
}

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (c *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	c.descriptorIndex = reader.readUint16()
}

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (c *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	c.bootstrapMethodAttrIndex = reader.readUint16()
	c.nameAndTypeIndex = reader.readUint16()
}
