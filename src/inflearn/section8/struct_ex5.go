// 구조체 심화(4)
package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

func (e Executives) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

type Executives struct {
	Employee     // is a 관계
	specialBonus float64
}

func main() {
	// 구조체 임베디드 메소드 오버라이딩 패턴

	// 예제1
	// 직원
	ep1 := Employee{"kim", 2000000, 150000}
	ep2 := Employee{"park", 15000000, 200000}

	// 임원
	ex := Executives{Employee{"lee", 5000000, 100000}, 100000}

	fmt.Println("ex1 : ", int(ep1.Calculate()))
	fmt.Println("ex1 : ", int(ep2.Calculate()))

	// Employee 구조체 통해서 메소드 호출
	fmt.Println("ex2 : ", int(ex.Calculate())) // 오버라이딩
	fmt.Println("ex2 : ", int(ex.Employee.Calculate()+ex.specialBonus))
}
