package chapter16

import (
	"sync/atomic"
	"fmt"
)

func AtomicTest()  {
	var a int32 = 10;
	atomic.SwapInt32(&a, 50)
	atomic.AddInt32(&a, 100)
	fmt.Println(a)
}