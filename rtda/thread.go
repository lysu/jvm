package rtda

import "github.com/lysu/jvm/rtda/heap"

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) GetFrames() []*Frame {
	return self.stack.getFrames()
}

func (self *Stack) isEmpty() bool {
	return self._top == nil
}

func (t *Thread) PC() int {
	return t.pc
}

func (t *Thread) SetPC(pc int) {
	t.pc = pc
}

func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}

func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}
func (t *Thread) TopFrame() *Frame {
	return t.CurrentFrame()
}

func (t *Thread) CurrentFrame() *Frame {
	return t.stack.top()
}

func (self *Thread) ClearStack() {
	self.stack.clear()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return NewFrame(self, method)
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}
