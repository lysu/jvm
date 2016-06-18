package constants

import (
	"github.com/lysu/jvm/instructions/base"
	"github.com/lysu/jvm/rtda"
)

type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
}
