package heap

type Object struct {
	class   *Class
	data    interface{} // Slots for Object, []int32 for int[] ...
	extra   interface{}
	monitor *Monitor
}

// create normal (non-array) object
func newObject(class *Class) *Object {
	return &Object{
		class:   class,
		data:    newSlots(class.instanceSlotCount),
		monitor: newMonitor(),
	}
}

// getters & setters
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.data.(Slots)
}
func (self *Object) Extra() interface{} {
	return self.extra
}
func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}
func (self *Object) Monitor() *Monitor {
	return self.monitor
}
func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

// reflection
func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}
func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}
