// 인터페이스 고급(1)
package main

import (
	"fmt"
)

func printContents(v interface{}) {
	fmt.Printf("Type : (%T)", v) // 원본파일
	fmt.Println("ex1 : ", v)
}

func main() {
	// 인터페이스 활용 (빈 인터페이스)
	// 빈 인터페이스 : 함수 매개변수, 리턴 값, 구조체 필드 등올 사용 가능 -> 어떤 타입으로도 변환 가능
	// 모든 타입을 나타내기 위해 빈 인터페이스 사용
	// 동적타입으로 생각하면 쉽다 (타 언어의 object 타입)

	// 예제1
	var a interface{}
	printContents(a)

	a = "Golang!"
	printContents(a)

	a = 7.5
	printContents(a)

	a = true
	printContents(a)

	a = nil
	printContents(a)

}
