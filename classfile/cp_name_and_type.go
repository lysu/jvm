package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex        uint16
	descriptionIndex uint16
}

func (c *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	c.nameIndex = reader.readUint16()
	c.descriptionIndex = reader.readUint16()
}
