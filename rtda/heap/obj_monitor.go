package heap

import (
	"sync"
)

type Monitor struct {
	owner      interface{} // *rtda.Thread
	ownerLock  sync.Locker
	lock       sync.Locker
	entryCount int
	cond       *sync.Cond
}

func newMonitor() *Monitor {
	m := &Monitor{}
	m.ownerLock = &sync.Mutex{}
	m.lock = &sync.Mutex{}
	m.cond = sync.NewCond(m.lock)
	return m
}

func (self *Monitor) Enter(thread interface{}) {
	defer self.ownerLock.Unlock()
	self.ownerLock.Lock()
	if self.owner == thread {
		self.entryCount++
		self.ownerLock.Unlock()
		return
	}
	self.lock.Lock()

	self.owner = thread
	self.entryCount = 1
}

func (self *Monitor) Exit(thread interface{}) {
	defer self.ownerLock.Unlock()
	self.ownerLock.Lock()
	var _unlock bool
	if self.owner == thread {
		self.entryCount--
		if self.entryCount == 0 {
			self.owner = nil
			_unlock = true
		}
	}

	if _unlock {
		self.lock.Unlock()
	}
}

func (self *Monitor) HasOwner(thread interface{}) bool {
	defer self.ownerLock.Unlock()
	self.ownerLock.Lock()
	isOwner := self.owner == thread

	return isOwner
}

func (self *Monitor) Wait() {
	defer self.ownerLock.Unlock()
	self.ownerLock.Lock()

	self.cond.Wait()
}

func (self *Monitor) NotifyAll() {
	self.cond.Broadcast()
}
