package control

import (
	"github.com/lysu/jvm/instructions/base"
	"github.com/lysu/jvm/rtda"
)

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
