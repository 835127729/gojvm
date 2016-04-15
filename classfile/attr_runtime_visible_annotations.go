package classfile

//RuntimeVisibleAnnotations
type RuntimeVisibleAnnotationsAttribute struct {
	annotation []Annotation
}

func (self *RuntimeVisibleAnnotationsAttribute) readInfo(reader *ClassReader) {
	numAnnotations := reader.readUint16()
	self.annotation = make([]Annotation, numAnnotations)
	for _, annotation := range self.annotation {
		annotation.typeIndex = reader.readUint16()
		numElementValuePairs := reader.readUint16()
		annotation.elementValuePairs = make([]ElementValuePair, numElementValuePairs)
		for _, elementValuePair := range annotation.elementValuePairs {
			elementValuePair.elementNameIndex = reader.readUint16()
			elementValuePair.value = readElementValue(reader)
		}
	}
}
