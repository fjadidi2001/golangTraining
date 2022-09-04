package main

import "fmt"

func main() {
	var (
		price1 float64
		price2 float64
	)
	fmt.Println("inter price in first year:")
	fmt.Scanln(&price1)
	fmt.Println("inter price in second year:")
	fmt.Scanln(&price2)
	inflation := (price2 - price1) / price1
	fmt.Println(inflation)
	g := inflation * price2
	price3 := g + price2
	fmt.Println(price3)

}
