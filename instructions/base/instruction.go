package base

import "github.com/lysu/jvm/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (i *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
}

type BranchInstruction struct {
	Offset int
}

func (i *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	i.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (i *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	i.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (i Index16Instruction) FetchOperands(reader *BytecodeReader) {
	i.Index = uint(reader.ReadUint16())
}
