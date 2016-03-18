package heap

type Object struct {
	class *Class
	data  interface{} // Slots for Object, []int32 for int[] ...
	extra interface{}
}

// create normal (non-array) object
func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}
