// 구조체 심화(1)
package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func NewAccount(number string, balance float64, interest float64) *Account { // 포인터 반환 아닌 경우 값 복사
	return &Account{number, balance, interest} // 구조체 인스턴스를 생성한 뒤 포인터 리턴
}

func main() {
	// 구조체 생성자 패턴 예제

	// 예제1
	park := NewAccount("245-903", 17000000, 0.04)
	fmt.Println("ex1 : ", park)

}
