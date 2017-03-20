package chapter16

import (
	"fmt"
	"time"
)

func sum(values []int, resultChan chan int)  {
	sum := 0
	for _,v:=range values {
		sum += v;
	}

	resultChan <- sum
	fmt.Println("Test", sum)
}

// 多个goroutine计算数值，通过通道传入主线程
func ChannelTest1()  {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int)
	go sum(values[0:1], resultChan)
    go sum(values[2:3], resultChan)
    go sum(values[3:4], resultChan)
    // sum1, sum2, sum3 := <-resultChan, <-resultChan, <-resultChan
    // fmt.Println("Result:", sum1, sum2, sum3)

	time.Sleep(1*time.Second)
	sum1, sum2 := <-resultChan, <-resultChan
	time.Sleep(1*time.Second)
	fmt.Println("Result:", sum1, sum2)

	time.Sleep(2*time.Second)
	fmt.Println("Main")
}