package main

import "fmt"

type person struct {
	fname string
	lname string
}

func (p person) fullname() {
	fmt.Printf("%s %s\n", p.fname, p.lname)

}
func main() {

	p := person{
		"fateme",
		"jadidi",
	}
	fmt.Println("welcome")

	defer p.fullname()

}
