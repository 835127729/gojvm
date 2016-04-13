package rtda

import (
	"fmt"
	"gojvm/rtda/heap"
)

/*
JVM
  Thread
    pc
    Stack
      Frame
        LocalVars
        OperandStack
*/
type Thread struct {
	pc    int // the address of the instruction currently being executed
	stack *Stack
	// todo
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) HandleUncaughtException(ex *heap.Object) {
	fmt.Println("*******************HandleUncaughtException****************************")
	self.stack.clear()
	sysClass := ex.Class().Loader().LoadClass("java/lang/System")
	sysErr := sysClass.GetStaticValue("out", "Ljava/io/PrintStream;").(*heap.Object)
	printStackTrace := ex.Class().GetInstanceMethod("printStackTrace", "(Ljava/io/PrintStream;)V")

	// call ex.printStackTrace(System.err)
	newFrame := self.NewFrame(printStackTrace)
	vars := newFrame.localVars
	vars.SetRef(0, ex)
	vars.SetRef(1, sysErr)
	self.PushFrame(newFrame)
}

//getter
func (self *Thread) PC() int {
	return self.pc
}
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}
func (self *Thread) GetFrames() []*Frame {
	return self.stack.getFrames()
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}
func (self *Thread) ClearStack() {
	self.stack.clear()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}
