package chapter16

import (
    "fmt"
    "sync"
	"time"
)

// 处于cond.Wait状态的所有gouroutine收到信号后将全部被唤醒并往下执行。需要注意的是，
// 从gouroutine执行完任务后，需要通过cond.L.Unlock释放锁， 否则其它被唤醒的gouroutine将没法继续执行
var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func test(x int) {

    cond.L.Lock() // 获取锁
    cond.Wait()   // 等待通知  暂时阻塞
    fmt.Println(x)
    time.Sleep(time.Second * 1)
    cond.L.Unlock() // 释放锁，不释放的话将只会有一次输出
}

func BroadcastTest()  {
	for i := 0; i < 40; i++ {
        go test(i)
    }
    fmt.Println("start all")
    cond.Broadcast() //  下发广播给所有等待的goroutine
    time.Sleep(time.Second * 60)
}



// cond实现同步
func SyncTest()  {
	locker := new(sync.Mutex)
    cond := sync.NewCond(locker)
    done := false

    cond.L.Lock()

    go func() {
        time.Sleep(2e9)
        done = true
        cond.Signal()
    }()

    if (!done) {
        cond.Wait()
    }

    fmt.Println("now done is ", done);
}

func CondTest()  {
	wait := sync.WaitGroup{}
	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)

	for i := 0; i < 3; i++ {
		go func(i int) {
			defer wait.Done()
			wait.Add(1)
			cond.L.Lock()
			fmt.Println("Waiting start...")
			cond.Wait()
			fmt.Println("Waiting end...")
			cond.L.Unlock()

			fmt.Println("Goroutine run. Number:", i)
		}(i)
	}

	// time.Sleep(1*time.Second)
	// cond.L.Lock()
	// cond.Signal()
	// cond.L.Unlock()

	// time.Sleep(1*time.Second)
	// cond.L.Lock()
	// cond.Signal()
	// cond.L.Unlock()

	// time.Sleep(1*time.Second)
	// cond.L.Lock()
	// cond.Signal()
	// cond.L.Unlock()

	time.Sleep(1*time.Second)
	cond.Broadcast();

	wait.Wait()
}


// 一个goroutine需要等待一批goroutine执行完毕以后才继续执行，那么这种多线程等待的问题就可以使用WaitGroup了。
// Add(n):添加或者减少等待goroutine的数量,n大于0表示添加n个等待goroutine的数量，小于0减少n个等待goroutine的数量
func WaitGroupTest()  {
	wp := new(sync.WaitGroup)
    wp.Add(10);

    for i := 0; i < 10; i++ {
		time.Sleep(1*time.Second)
        go func() {
            fmt.Println("done ", i)
            wp.Done()
        }()
    }

    wp.Wait()
    fmt.Println("wait end")
}