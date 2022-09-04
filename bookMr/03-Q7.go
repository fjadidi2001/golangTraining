package main

import "fmt"

func main() {
	var (
		salary    float64
		Insurance float64
		tax       float64
	)
	fmt.Println("enter your salary:")
	fmt.Scanln(&salary)
	Insurance = 0.1 * salary
	tax = 0.07 * salary
	salary = salary - tax - Insurance
	fmt.Println(salary)
}
