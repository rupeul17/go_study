// switch문(3)
package main

import "fmt"

func main() {
  // 예제1
  a := 30 / 15
  switch a {
  case 2, 4, 6:
    fmt.Println("a -> ", a, "는 짝수")
  case 1, 3, 5:
    fmt.Println("a -> ", a , "는 홀수")
  }

  // 예제2
  // go 언어는 break가 자동으로 됨
  // fallthrough 쓰면 해당 case의 밑 case 소스도 실행
  switch e := "go"; e {
  case "java":
    fmt.Println("java!")
    fallthrough
  case "go":
    fmt.Println("go!")
    fallthrough
  case "python":
    fmt.Println("python!")
  case "ruby":
    fmt.Println("ruby!")
  case "c":
    fmt.Println("c!")
  }
}
