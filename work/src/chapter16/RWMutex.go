package chapter16

import (
	"time"
	"sync"
)

var m *sync.RWMutex
var val = 0

func read(i int) {
    // m.RLock()
    time.Sleep(1 * time.Second)
    println("val: ", val)
    time.Sleep(1 * time.Second)
    // m.RUnlock()
}

func write(i int) {
	m.Lock()
    val = 10
	time.Sleep(1 * time.Second)
	m.Unlock()
}

func RWMutexTest()  {
	m = new(sync.RWMutex)
    go read(1)
    go write(2)
    go read(3)
    time.Sleep(5 * time.Second)
}
