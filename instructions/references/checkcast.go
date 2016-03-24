package references

import "gojvm/instructions/base"
import "gojvm/rtda"
import "gojvm/rtda/heap"

// Check whether object is of given type
type CHECK_CAST struct{ base.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ConstantClassref)
	class := classRef.ResolveClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}