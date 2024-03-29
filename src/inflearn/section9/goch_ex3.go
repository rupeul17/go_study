// 채널 심화(3)
package main

import (
	"fmt"
)

func recvOnly(cnt int) <-chan int {
	sum := 1
	tot := make(chan int)
	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		tot <- sum
		tot <- 777
		tot <- 777
		close(tot)
	}()

	return tot
}

func total(c <-chan int) <-chan int { // 리턴 타입은 수신 전용 채널
	tot := make(chan int)
	go func() {
		a := 0
		for i := range c {
			a = a + i
		}
		tot <- a
	}()
	return tot
}

func main() {
	// 채널

	// 예제1
	c := recvOnly(100) // 채널 반환
	output := total(c) // 채널 전달 후 반환

	//output <- 777	// 예외
	fmt.Println("ex1 : ", <-output)
}
