package main

import "fmt"

var sima struct {
	fname string
	lname string
	age   int
	score float64
}



func main() {
	sima.fname = "fateme"
	sima.lname = "Jadidi"
	sima.age = 21
	sima.score = 16.79
	fmt.Println(sima)
}
