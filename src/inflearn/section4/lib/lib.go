// 라이브러리 접근제어(1)
package lib // 패키지 이름은 폴더명

import "fmt"

func init() {
	fmt.Println("lib package > init start!")
}

func CheckNum(c int32) bool {
	return c > 10
}
