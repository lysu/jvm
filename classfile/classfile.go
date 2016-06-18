package classfile

import "fmt"

type ClassFile struct {
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", err)
			}
		}
	}()
	cr := &ClassReader{data: classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (c *ClassFile) MinorVersion() uint16 {
	return c.minorVersion
}

func (c *ClassFile) MajorVersion() uint16 {
	return c.majorVersion
}

func (c *ClassFile) ConstantPool() ConstantPool {
	return c.constantPool
}

func (c *ClassFile) AccessFlags() uint16 {
	return c.accessFlags
}

func (c *ClassFile) Fields() []*MemberInfo {
	return c.fields
}

func (c *ClassFile) Methods() []*MemberInfo {
	return c.methods
}

func (c *ClassFile) ClassName() string {
	return c.constantPool.getClassName(c.thisClass)
}

func (c *ClassFile) SuperClassName() string {
	if c.superClass > 0 {
		return c.constantPool.getClassName(c.superClass)
	}
	return ""
}

func (c *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(c.interfaces))
	for i, cpIndex := range c.interfaces {
		interfaceNames[i] = c.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

func (c *ClassFile) read(cr *ClassReader) {
	c.readAndCheckMagic(cr)
	c.readAndCheckVersion(cr)
	c.constantPool = readConstantPool(cr)
	c.accessFlags = cr.readUint16()
	c.thisClass = cr.readUint16()
	c.superClass = cr.readUint16()
	c.interfaces = cr.readUint16s()
	c.fields = readMembers(cr, c.constantPool)
	c.methods = readMembers(cr, c.constantPool)
	c.attributes = readAttributes(cr, c.constantPool)
}

func (c *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (c *ClassFile) readAndCheckVersion(reader *ClassReader) {
	c.minorVersion = reader.readUint16()
	c.majorVersion = reader.readUint16()
	switch c.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if c.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

