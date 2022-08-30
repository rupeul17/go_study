// 채널 심화(4)
package main

import (
	"fmt"
	"time"
)

func main() {
	// 채널 셀렉트 구문
	// 채널 select 구문 -> 채널에 값이 수신되면 해당 case문 실행
	// 일회성 구문이므로 for문 안에서 수행
	// default 구문 처리 주의

	// 예제1
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- 77
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "golang hi!"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("ch1 : ", num)
			case str := <-ch2:
				fmt.Println("ch1 : ", str)
				// default : 수발신 용으로 쓸 때는 default 사용 안함
			}
		}
	}()

	time.Sleep(7 * time.Second)
}
