package alo

import (
	"errors"
	"runtime"
	"sync/atomic"
)

const lockOff = 0
const lockOn = 1

// AtomicLock is mutual exclusive lock
type AtomicLock struct {
	locked int32
}

// Lock will set `locked` to 1
func (al *AtomicLock) Lock() {
	for !atomic.CompareAndSwapInt32(&al.locked, lockOff, lockOn) {
		runtime.Gosched()
	}
}

// Unlock will set `locked` to 0
func (al *AtomicLock) Unlock() {
	if al.locked == lockOff {
		panic(errors.New("alo.Unlock: the lock has already unlocked"))
	}
	for !atomic.CompareAndSwapInt32(&al.locked, lockOn, lockOff) {
		runtime.Gosched()
	}
}
