package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (c ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := c[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index")
}

func (c ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := c.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := c.getUtf8(ntInfo.nameIndex)
	typ := c.getUtf8(ntInfo.descriptionIndex)
	return name, typ
}

func (c ConstantPool) getClassName(index uint16) string {
	classInfo := c.getConstantInfo(index).(*ConstantClassInfo)
	return c.getUtf8(classInfo.nameIndex)
}

func (p ConstantPool) getUtf8(cpIndex uint16) string {
	utf8Info := p.getConstantInfo(cpIndex).(*ConstantUtf8Info)
	return utf8Info.str
}
