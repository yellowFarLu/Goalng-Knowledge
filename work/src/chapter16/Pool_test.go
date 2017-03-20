package chapter16

/*
线程池的使用及压力测试
*/

import (
	"sync"
	"testing"
)

var bytePool = sync.Pool {
	New : newPool,	// 赋值一个函数指针常量（这里是函数名字）
}

// 任何对象都是实现interface{}，因此空接口的变量能接收任何值
func newPool() interface{} {
	b := make([]byte, 1024)
	return &b;
}

// 性能测试
// 测试对象池
func BenchmarkPool(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		// New函数return什么，Get()后面就跟什么类型。比如New返回【】byte的地址，Get（）后面跟*[]byte
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj);
	}
}

// 测试不用对象池
func BenchmarkAlloc(b *testing.B) {
    for i := 0; i < b.N; i++ {
        obj := make([]byte, 1024)
        _ = obj
    }
}
