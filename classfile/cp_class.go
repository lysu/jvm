package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (c *ConstantClassInfo) readInfo(reader *ClassReader) {
	c.nameIndex = reader.readUint16()
}

func (c *ConstantClassInfo) Name() string {
	return c.cp.getUtf8(c.nameIndex)
}
