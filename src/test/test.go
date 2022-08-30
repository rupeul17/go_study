package main

import "fmt"

type Job struct {
	name   string
	number int
}

func main() {

	var num int = 5

	//var job [num]Job

	job := make([]Job, num)

	//job = append(job)

	job[0].name = "test"
	job[0].number = 1

	fmt.Println(job)

}
