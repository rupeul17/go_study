// 자료형 : 맵(4)
package main

import "fmt"

func main() {
	// 맵 (map)
	// 맵 조회 할 경우에 주의할 점

	// 예제1
	map1 := map[string]int{
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1 := map1["lemon"]
	value2, ok1 := map1["lemon"] // 값 0이 출력
	value3, ok2 := map1["kiwi"]  // 없을시 기본값 0이 출력

	fmt.Println("ex1 : ", value1, value2, ok1, value3, ok2) // ok로 키의 유무 확인 (true, false)

	// 예제2
	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : kiwi is not exist!")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("ex2 : kiwi is not exist!")
	}
}
