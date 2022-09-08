package main

import "fmt"

func main() {
	// 짧은 선언
	// 함수 안에서만 사용(전역x), 선언 후 할당 예외 발생
	// 주로 제한된 범위의 함수내에서 사용할 경우 코드 가독성을 높일 수 있다 -> 추천
	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	//shortVar3 := true // 예외 발생

	fmt.Println("sVar1 : ", shortVar1, "\n", "sVar2 : ", shortVar2, "\n", "sVar3 : ", shortVar3)

}
