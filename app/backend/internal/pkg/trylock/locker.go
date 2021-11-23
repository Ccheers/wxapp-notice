package trylock

import (
	"errors"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

var (
	errGetLockTimeOut = errors.New("get lock timeout")
)

func IsErrGetLockTimeOut(err error) bool {
	return errors.Is(err, errGetLockTimeOut)
}

type TryMutexLocker interface {
	sync.Locker
	TryLock(duration time.Duration) error
}
type atomicMutex struct {
	mu sync.Mutex
}

func NewAtomicMutex() TryMutexLocker {
	return &atomicMutex{}
}

const (
	UnLock = 0
	Locked = 1
)

func (a *atomicMutex) TryLock(duration time.Duration) error {
	if duration > 0 {
		timeoutCh := time.After(duration)
		for {
			select {
			case <-timeoutCh:
				return errGetLockTimeOut
			default:
				// 这里原子操作 sync.mutex 的 state 指针 将其尝试 置为 1
				if atomic.CompareAndSwapInt64((*int64)(unsafe.Pointer(&a.mu)), UnLock, Locked) {
					return nil
				} else {
					runtime.Gosched()
				}
			}
		}
	}

	switch {
	case atomic.CompareAndSwapInt64((*int64)(unsafe.Pointer(&a.mu)), UnLock, Locked):
		return nil
	default:
		return errGetLockTimeOut
	}
}

func (a *atomicMutex) Lock() {
	a.mu.Lock()
}

func (a *atomicMutex) Unlock() {
	a.mu.Unlock()
}
