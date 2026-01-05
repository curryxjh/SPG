package sync

import (
	gosync "sync"
)

// Locker defines a abstract interface for lockers.
type Locker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

var _ Locker = (*gosync.RWMutex)(nil)

type FakeLocker struct {
}

func (f FakeLocker) Lock() {
}

func (f FakeLocker) Unlock() {
}

func (f FakeLocker) RLock() {
}

func (f FakeLocker) RUnlock() {
}
