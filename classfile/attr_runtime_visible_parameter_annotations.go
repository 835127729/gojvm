package classfile

type RuntimeVisibleParammeterAnnotationsAttribute struct {
	parameterAnnotations []ParameterAnnotations
}

func (self *RuntimeVisibleParammeterAnnotationsAttribute) readInfo(reader *ClassReader) {
	numParameters := reader.readUint8()
	self.parameterAnnotations = make([]ParameterAnnotations, numParameters)
	for i, _ := range self.parameterAnnotations {
		numAnnotations := reader.readUint16()
		self.parameterAnnotations[i] = make([]Annotation, numAnnotations)
		for _, annotation := range self.parameterAnnotations[i] {
			annotation.typeIndex = reader.readUint16()
			numElementValuePairs := reader.readUint16()
			annotation.elementValuePairs = make([]ElementValuePair, numElementValuePairs)
			for _, elementValuePair := range annotation.elementValuePairs {
				elementValuePair.elementNameIndex = reader.readUint16()
				elementValuePair.value = readElementValue(reader)
			}
		}
	}
}

type ParameterAnnotations []Annotation
