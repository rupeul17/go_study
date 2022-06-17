// 자료형 : 맵(3)
package main

import "fmt"

func main() {
	// 맵 (MAP)
	// 맵 값 변경 및 삭제

	// 예제1
	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
		"home1":  "http://test1.com",
	}

	fmt.Println("ex1 : ", map1)

	map1["home2"] = "http://test2.com" // 추가
	fmt.Println("ex1 : ", map1)

	map1["home2"] = "http://test2_2.com" // 변경
	fmt.Println("ex1 : ", map1)

	delete(map1, "home2") // 삭제
	fmt.Println("ex1 : ", map1)
}
