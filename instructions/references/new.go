package references

import (
	"github.com/lysu/jvm/instructions/base"
	"github.com/lysu/jvm/rtda"
	"github.com/lysu/jvm/rtda/heap"
)

type NEW struct{ base.Index16Instruction }

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
