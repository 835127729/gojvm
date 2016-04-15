package classfile

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/
type SourceDebugExtensionAttribute struct {
	debugExtension []uint8
}

func (self *SourceDebugExtensionAttribute) readInfo(reader *ClassReader) {
	attributeLength := reader.readUint32()
	self.debugExtension = make([]uint8, attributeLength)
}

func (self *SourceDebugExtensionAttribute) DebugExtension() []uint8 {
	return self.debugExtension
}
