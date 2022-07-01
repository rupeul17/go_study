// 함수 defer(4)
package main

import (
	"fmt"
)

func start(t string) string {
	fmt.Println("start : ", t)
	return t
}

func end(t string) {
	fmt.Println("end : ", t)
}

func a() {
	defer end(start("b")) // defer는 바로 뒤 end 함수만 지연함수로 걸음
	fmt.Println("in a")
}

func main() {
	// 예제1
	a()
}
