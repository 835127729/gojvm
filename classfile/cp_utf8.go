package classfile

import "fmt"
import "unicode/utf16"

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

func (self *ConstantUtf8Info) Str() string {
	return self.str
}

func (self *ConstantUtf8Info) toString() string {
	return "Utf8"
}

/*
func decodeMUTF8(bytes []byte) string {
	return string(bytes) // not correct!
}
*/

// mutf8 -> utf16 -> utf32 -> string
// see java.io.DataInputStream.readUTF(DataInput)
func decodeMUTF8(bytearr []byte) string {
	utflen := len(bytearr)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	for count < utflen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararr_count] = c&0x1F<<6 | char2&0x3F
			chararr_count++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-2])
			char3 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararr_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararr_count++
		default:
			/* 10xx xxxx,  1111 xxxx */
			count += 6
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char5 := uint16(bytearr[count-5])
			char4 := uint16(bytearr[count-4])
			char3 := uint16(bytearr[count-3])
			char2 := uint16(bytearr[count-2])
			char1 := uint16(bytearr[count-1])
			if c != 0xED || char1&0xF0 != 0xA0 || char2&0xC0 != 0x80 || char3 != 0xED || char4&0xF0 != 0xB0 || char5&0x60 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararr_count] = 0x0001
			chararr[chararr_count+1] = 0x0000 | ((char1 & 0x0F) << 16) | ((char2 & 0x3F) << 10) | ((char4 & 0x0F) << 6) | (char5 & 0x3f)
			chararr_count += 2
		}
	}
	// The number of chars produced may be less than utflen
	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes)
}
