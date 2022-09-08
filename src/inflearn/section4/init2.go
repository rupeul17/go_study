// go 초기화 함수(1)
package main

import (
	"fmt"
	"section4/lib" // import 하는 시점에서 패키지 내 init() 호출
)

var num int32

// 변수 초기화
func init() {
	num = 30
	fmt.Println("Main Init Start!")
}

func main() {
	fmt.Println("10보다 큰수? : ", lib.CheckNum(num))
}
