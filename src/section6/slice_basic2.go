// 자료형 : 슬라이스(2)
package main

import "fmt"

func main() {
	// 슬라이스 (슬라이스 참조 타입 증명)

	// 예제1 (배열)
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1 // 복사
	arr2[0] = 7

	fmt.Println("ex1 : ", arr1)
	fmt.Println("ex1 : ", arr2)
	fmt.Println()

	// 결론 : 배열은 값 복사

	// 예제2 (슬라이스)
	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = slice1
	slice2[0] = 7

	fmt.Println("ex2 : ", slice1)
	fmt.Println("ex2 : ", slice2)
	fmt.Println()

	// 결론 : 슬라이스는 참조 처리

	// 예제3 (슬라이스 예외 상황)
	slice3 := make([]int, 50, 100) // 50(길이)만큼 0으로 초기화
	fmt.Println("ex3 : ", slice3[4])
	//fmt.Println("ex3 : ", slice3[50])	// 길이 index over 예외처리
	fmt.Println()

	// 예제4 (슬라이스 순회)
	for i, v := range slice1 {
		fmt.Println("ex4 : ", i, v)
	}
}
