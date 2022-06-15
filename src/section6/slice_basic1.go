// 자료형 : 슬라이스(1)
package main

import "fmt"

func main() {
	// 길이x(가변) -> 동적으로 크기가 늘어난다, 레퍼런스(참조 값) 타입
	// 슬라이스 (길이 & 용량) 크기가 동적으로 할당 가능

	// 2가지 선언 방법 1. 배열처럼 선언, 2. make 함수 : make(자료형, 길이, 용량(생략시 길이))

	// 예제1
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	//slice2[0]	// 값 수정 불가능 -> 초기화가 안되어 있음
	slice3[4] = 10 // 값 수정 가능 -> 이미 초기화가 되어 있음

	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)

	// 초기화 안해줌

	// 예제2
	// 용량을 넘어선 자료가 들어오면 메모리 재할당 -> 성능 저하
	var slice5 []int = make([]int, 5, 10)
	var slice6 = make([]int, 5)
	slice7 := make([]int, 5, 100)
	slice8 := make([]int, 5)

	slice6[2] = 7 // 삽입 -> 초기화 안되어 있어도 값 변경 가능

	fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("%-5T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)

	// make 로 초기화 시 0으로 초기화 해줌

	// 예제3
	var slice9 []int // nil 슬라이스 (길이와 용량 0)
	if slice9 == nil {
		fmt.Println("This is Nil Slice")
	}
}
