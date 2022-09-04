package main

import "fmt"

func main() {
	name := "fateme"
	fmt.Printf("orginal string :%s\n", string(name))
	fmt.Printf("reversed string :")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}
