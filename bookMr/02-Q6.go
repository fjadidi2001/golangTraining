package main

import "fmt"

func main() {
	var (
		yearInSecond = 3.156 * 10e7
		age          float64
	)
	fmt.Println("enter your age plz:")
	fmt.Scanln(&age)
	ageInsecond := age * yearInSecond
	fmt.Println(ageInsecond)

}
