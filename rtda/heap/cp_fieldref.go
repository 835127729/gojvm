package heap

import (
	"fmt"
	cf "gojvm/classfile"
	"gojvm/jutil"
)

type ConstantFieldref struct {
	ConstantMemberref
	field *Field
}

func newConstantFieldref(refInfo *cf.ConstantFieldrefInfo) *ConstantFieldref {
	ref := &ConstantFieldref{}
	ref.copy(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *ConstantFieldref) String() string {
	return fmt.Sprintf("{ConstantFieldref className:%v name:%v descriptor:%v}",
		self.className, self.name, self.descriptor)
}

func (self *ConstantFieldref) InstanceField() *Field {
	if self.field == nil {
		self.resolveInstanceField()
	}
	return self.field
}

func (self *ConstantFieldref) resolveInstanceField() {
	fromClass := bootLoader.LoadClass(self.className)

	field := fromClass.getField(self.name, self.descriptor, false)
	if field != nil {
		self.field = field
		return
	}

	// todo
	jutil.Panicf("instance field not found! %v", self)
}

func (self *ConstantFieldref) StaticField() *Field {
	if self.field == nil {
		self.resolveStaticField()
	}
	return self.field
}

func (self *ConstantFieldref) resolveStaticField() {
	fromClass := bootLoader.LoadClass(self.className)

	field := fromClass.getField(self.name, self.descriptor, true)
	if field != nil {
		self.field = field
		return
	}

	// todo
	jutil.Panicf("instance field not found! %v", self)
}
