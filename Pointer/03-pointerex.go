package main

import "fmt"

func main() {
	i, j := 45, 12
	fmt.Println(i, j)
	fmt.Println(&i, &j) //print the address of i and the address of j
}
