package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt") //os package:func Open(name string) (*File, error)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}
