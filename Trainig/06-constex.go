package main

import "fmt"

const (
	add = iota
	subtract
	multiply
	divide
)

type operation int

func (op operation) calculate(lhs, rhs int) int {
	switch op {
	case add:
		return lhs + rhs
	case subtract:
		return lhs - rhs
	case multiply:
		return lhs * rhs
	case divide:
		return lhs / rhs

	}
	panic("unhandled operation")
}

func main() {
	Add := operation(add)
	fmt.Println(add.calculate(2, 2))
	Subtract := operation(subtract)

	fmt.Println(subtract.calculate(2, 2))
	Multiply := operation(multiply)

	fmt.Println(multiply.calculate(2, 2))
	diivide := operation(divide)

	fmt.Println(divide.calculate(2, 2))
}
