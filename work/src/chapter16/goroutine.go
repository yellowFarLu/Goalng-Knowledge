package chapter16
import "fmt"
import "time"
import "runtime"

func loop(t int) {
    fmt.Println("loop starts!")
    for i := 0; i < 1000000; i++ {
        fmt.Printf("%d ---- %d\n", i, t)
    }
    fmt.Println()
}

func GoroutineTest1() {
    runtime.GOMAXPROCS(2) // 强制使用1个cpu

    go loop(-1)
    go loop(-2)

    time.Sleep(3 * time.Second)
}


func Add(x, y int) {
    z := x + y
    fmt.Println(z)
}


func GoroutineTest2()  {
	// runtime.GOMAXPROCS(1) // 强制使用1个cpu

    go Add(0, 0)
    go Add(1, 1)
    go Add(2, 2)
    go Add(3, 3)
    go Add(4, 4)
    go Add(5, 5)
    go Add(6, 6)
    go Add(7, 7)
    go Add(8, 8)
    go Add(9, 9)

    fmt.Println("main goroutine finished!")
    time.Sleep(3 * time.Second)
}