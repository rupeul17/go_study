// 구조체 심화(4)
package main

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

type Executives struct {
	Employee
	specialBonus float64
}

func main() {

}
