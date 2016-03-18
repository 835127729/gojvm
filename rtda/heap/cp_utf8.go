package heap

import (
	cf "gojvm/classfile"
)

type ConstantUtf8 struct {
	str string
}

func newConstantUtf8(utf8Info *cf.ConstantUtf8Info) *ConstantUtf8 {
	return &ConstantUtf8{
		str: utf8Info.Str(),
	}
}

func (self *ConstantUtf8) Str() string {
	return self.str
}
