// 데이터 타입 : numeric 연산(2)
package main

import (
	"fmt"
)

func main() {
	// 예제1
	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println("ex1 : ", n1+n2)
	fmt.Println("ex1 : ", n1-n2)
	fmt.Println("ex1 : ", n1*n2)
	fmt.Println("ex1 : ", n1/n2)
	fmt.Println("ex1 : ", n1%n2)
	fmt.Println("ex1 : ", n1<<2)
	fmt.Println("ex1 : ", n1>>2)
	fmt.Println("ex1 : ", ^n1) // 반전 연산자 (0->1, 1->0)

	// 예제2
	var n3 int = 12
	var n4 float32 = 8.2
	var n5 uint16 = 1024
	var n6 uint32 = 120000

	fmt.Println("ex3 : ", float32(n3)+n4) // 정수 + 실수 계산 시 보통 실수형에 맞춰서 계산
	fmt.Println("ex3 : ", n3+int(n4))     // 정수 + 실수 계산 시 보통 실수형에 맞춰서 계산
	fmt.Println("ex3 : ", n5+uint16(n6))  // max 값을 넘으면 강제로 max값에 맞추고 계산
}
