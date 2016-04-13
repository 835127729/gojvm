package references

import (
	"gojvm/instructions/base"
	"gojvm/rtda"
)

// Enter monitor for object
type MONITOR_ENTER struct{ base.NoOperandsInstruction }

func (self *MONITOR_ENTER) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		panic("NullPointerException")
	} else {
		ref.Monitor().Enter(thread)
	}
}

// Exit monitor for object
type MONITOR_EXIT struct{ base.NoOperandsInstruction }

func (self *MONITOR_EXIT) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		frame.RevertNextPC()
		panic("NullPointerException")
	} else {
		ref.Monitor().Exit(thread)
	}
}
