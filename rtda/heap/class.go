package heap

import "strings"

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
}

func (c *Class) InitStarted() bool {
	return c.initStarted
}

func (c *Class) StartInit() {
	c.initStarted = true
}

func (c *Class) SuperClass() *Class {
	return c.superClass
}

func (c *Class) Name() string {
	return c.name
}

func (c *Class) StaticVars() Slots {
	return c.staticVars
}

func (c *Class) isAccessibleTo(other *Class) bool {
	return c.IsPublic() || c.GetPackageName() == other.GetPackageName()
}

func (c *Class) GetPackageName() string {
	if i := strings.LastIndex(c.name, "/"); i >= 0 {
		return c.name[:i]
	}
	return ""
}

func (c *Class) ConstantPool() *ConstantPool {
	return c.constantPool
}

func (c *Class) NewObject() *Object {
	return newObject(c)
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}
