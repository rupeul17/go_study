// 데이터 타입 : 문자열 연산(2)
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 예제(1) (결합 : 일반연산)
	str1 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb" +
		"ccccccccccccccccccccccccccccccccc"

	str2 := "ddddddddddddddddddddddddddddddddd"

	fmt.Println("ex1 : ", str1+str2)

	// 예제(2) (결합 : join)
	strSet := []string{}          // 슬라이스 선언
	strSet = append(strSet, str1) // 뒤에 붙이는 함수
	strSet = append(strSet, str2)

	fmt.Println("ex2 : ", strings.Join(strSet, "-----"))
}
