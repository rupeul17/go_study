// 패키지 접근제어(2)
package main

import (
	"fmt"
	checkUp "section4/lib"
	_ "section4/lib2" // 밑줄 처리 -> 저장 시 사용하지 않는 패키지 자동삭제 -> 밑줄 처리하면 사용 안해도 일단 pass
)

func main() {
	// 패키지 접근제어
	// 별칭 사용
	// 빈 식별자 사용

	fmt.Println("10보다 큰 수? : ", checkUp.CheckNum(11))
}
