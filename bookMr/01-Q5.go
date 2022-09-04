package main

import (
	"fmt"
)

func main() {
	var m, i = 3 * 10e23, float64(950)
	var w float64
	fmt.Println(m, i)
	fmt.Println("input w")
	fmt.Scanln(&w)
	count := (m * i) / w
	fmt.Println(count)
}
