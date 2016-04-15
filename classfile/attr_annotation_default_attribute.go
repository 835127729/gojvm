package classfile

type AnnotationDefaultAttribute struct {
	defaultValue ElementValue
}

func (self *AnnotationDefaultAttribute) readInfo(reader *ClassReader) {
	self.defaultValue = readElementValue(reader)
}
