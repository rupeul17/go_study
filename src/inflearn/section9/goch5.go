// 채널(channel) 기초(5)
package main

import (
	"fmt"
)

func main() {
	// 채널
	// close : 채널 닫기, 주의 -> 닫힌 채널에 값 전송 시 패닉(예외) 발생
	// range : 채널 안에서 값을 꺼낸다. (순회), 채널 닫아야 (close) 반복문 종료 -> 채널이 열려 있고 값 전송하지 않으면 계속 대기!

	// 예제1
	ch := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- true
		}
		close(ch) // 5회 채널에 값 전송 후 채널 닫기
	}()

	for i := range ch { // 채널에서 값을 꺼내온다(채널이 close 될때까지)
		fmt.Println("ex1 : ", i)
	}
}
