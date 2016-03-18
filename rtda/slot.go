package rtda

import "gojvm/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
