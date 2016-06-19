package heap

type Object struct {
	class  *Class
	fields Slots
}

func (o *Object) Fields() Slots {
	return o.fields
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}
