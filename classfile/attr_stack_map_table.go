package classfile

import (
	"fmt"
)

/*
StackMapTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
*/
type StackMapTableAttribute struct {
	cp              ConstantPool
	attributeLength uint32
	entries         []StackMapFrame
}

//VerificationTypeInfo Stack
type VerificationTypeInfo interface {
	name() string
}

const (
	ITEM_Top               uint8 = 0
	ITEM_Integer           uint8 = 1
	ITEM_Float             uint8 = 2
	ITEM_Long              uint8 = 3
	ITEM_Double            uint8 = 4
	ITEM_Null              uint8 = 5
	ITEM_UninitializedThis uint8 = 6
	ITEM_Object            uint8 = 7
	ITEM_Uninitialized     uint8 = 8
)

type TopVariableInfo struct {
}

func (self TopVariableInfo) name() string {
	return "top"
}

type IntegerVariableInfo struct {
}

func (self IntegerVariableInfo) name() string {
	return "int"
}

type FloatVariableInfo struct {
}

func (self FloatVariableInfo) name() string {
	return "float"
}

type LongVariableInfo struct {
}

func (self LongVariableInfo) name() string {
	return "long"
}

type DoubleVariableInfo struct {
}

func (self DoubleVariableInfo) name() string {
	return "double"
}

type NullVariableInfo struct {
}

func (self NullVariableInfo) name() string {
	return "Null"
}

type UninitializedThisVariableInfo struct {
}

func (self UninitializedThisVariableInfo) name() string {
	return "UninitializedThis"
}

type ObjectVariableInfo struct {
	cp         ConstantPool
	cpoolIndex uint16
}

func (self ObjectVariableInfo) name() string {
	return "class " + self.cp.getClassName(self.cpoolIndex)
}

type UninitializedVariableInfo struct {
	offset uint16
}

func (self UninitializedVariableInfo) name() string {
	return "Uninitialized"
}

//帧类型
type StackMapFrame interface {
	offset() uint16
}

type SameFrame struct {
	StackMapFrame
	frameType uint8
}

func (self *SameFrame) offset() uint16 {
	return uint16(self.frameType)
}

type SameLocals1StackItemFrame struct {
	StackMapFrame
	frameType uint8
	stack     []VerificationTypeInfo
}

func (self *SameLocals1StackItemFrame) offset() uint16 {
	return uint16(self.frameType - 64)
}

type SameLocals1StackItemFrameExtended struct {
	StackMapFrame
	frameType   uint8
	offsetDelta uint16
	stack       []VerificationTypeInfo
}

func (self *SameLocals1StackItemFrameExtended) offset() uint16 {
	return self.offsetDelta
}

type ChopFrame struct {
	StackMapFrame
	frameType   uint8
	offsetDelta uint16
}

func (self *ChopFrame) offset() uint16 {
	return self.offsetDelta
}

type SameFrameExtended struct {
	StackMapFrame
	frameType   uint8
	offsetDelta uint16
}

func (self *SameFrameExtended) offset() uint16 {
	return self.offsetDelta
}

type AppendFrame struct {
	StackMapFrame
	frameType   uint8
	offsetDelta uint16
	locals      []VerificationTypeInfo
}

func (self *AppendFrame) offset() uint16 {
	return self.offsetDelta
}

type FullFrame struct {
	StackMapFrame
	frameType   uint8
	offsetDelta uint16
	stack       []VerificationTypeInfo
	locals      []VerificationTypeInfo
}

func (self *FullFrame) offset() uint16 {
	return self.offsetDelta
}

func (self *StackMapTableAttribute) readInfo(reader *ClassReader) {
	numberOfEnties := reader.readUint16()
	self.entries = make([]StackMapFrame, numberOfEnties)
	for i := range self.entries {
		self.entries[i] = self.readStackMapFrame(reader)
	}
}

func (self *StackMapTableAttribute) readStackMapFrame(reader *ClassReader) StackMapFrame {
	tag := reader.readUint8()
	fmt.Println(tag)
	if tag >= 0 && tag < 64 {
		return &SameFrame{frameType: tag}
	} else if tag >= 64 && tag < 128 {
		sameLocals1StackItemFrame := &SameLocals1StackItemFrame{
			frameType: tag,
		}
		sameLocals1StackItemFrame.stack = make([]VerificationTypeInfo, 1)
		sameLocals1StackItemFrame.stack[0] = self.readVerificationTypeInfo(reader)
		return sameLocals1StackItemFrame
	} else if tag == 247 {
		sameLocals1StackItemFrameExtended := &SameLocals1StackItemFrameExtended{
			frameType: tag,
		}
		sameLocals1StackItemFrameExtended.offsetDelta = reader.readUint16()
		sameLocals1StackItemFrameExtended.stack = make([]VerificationTypeInfo, 1)
		sameLocals1StackItemFrameExtended.stack[0] = self.readVerificationTypeInfo(reader)
		return sameLocals1StackItemFrameExtended
	} else if tag > 247 && tag < 251 {
		chopFrame := &ChopFrame{
			frameType:   tag,
			offsetDelta: reader.readUint16(),
		}
		return chopFrame
	} else if tag == 251 {
		sameFrameExtended := &SameFrameExtended{
			frameType: tag,
		}
		sameFrameExtended.offsetDelta = reader.readUint16()
		return sameFrameExtended
	} else if tag >= 252 && tag <= 254 {
		appendFrame := &AppendFrame{
			frameType: tag,
		}
		appendFrame.offsetDelta = reader.readUint16()
		appendFrame.locals = make([]VerificationTypeInfo, appendFrame.frameType-251)
		count := len(appendFrame.locals)
		for i := 0; i < count; i++ {
			appendFrame.locals[i] = self.readVerificationTypeInfo(reader)
		}
		return appendFrame
	} else if tag == 255 {
		fullFrame := &FullFrame{
			frameType: tag,
		}
		fullFrame.offsetDelta = reader.readUint16()
		count := reader.readUint16()
		fullFrame.locals = make([]VerificationTypeInfo, count)
		for i := range fullFrame.locals {
			fullFrame.locals[i] = self.readVerificationTypeInfo(reader)
		}
		count = reader.readUint16()
		fullFrame.stack = make([]VerificationTypeInfo, count)
		for i := range fullFrame.stack {
			fullFrame.stack[i] = self.readVerificationTypeInfo(reader)
		}
		return fullFrame
	} else {
		panic("stack map table error")
	}
}

func (self *StackMapTableAttribute) readVerificationTypeInfo(reader *ClassReader) VerificationTypeInfo {
	tag := reader.readUint8()
	switch tag {
	case ITEM_Top:
		return &TopVariableInfo{}
	case ITEM_Integer:
		return &IntegerVariableInfo{}
	case ITEM_Float:
		return &FloatVariableInfo{}
	case ITEM_Long:
		return &LongVariableInfo{}
	case ITEM_Double:
		return &DoubleVariableInfo{}
	case ITEM_Null:
		return &NullVariableInfo{}
	case ITEM_UninitializedThis:
		return &UninitializedThisVariableInfo{}
	case ITEM_Object:
		objectVariableInfo := ObjectVariableInfo{cp: self.cp}
		objectVariableInfo.cpoolIndex = reader.readUint16()
		return &objectVariableInfo
	case ITEM_Uninitialized:
		uninitializedVariableInfo := UninitializedVariableInfo{}
		uninitializedVariableInfo.offset = reader.readUint16()
		return &uninitializedVariableInfo
	default:
		panic("stackmaptable error")
	}
}
