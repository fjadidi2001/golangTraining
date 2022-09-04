package main

import "fmt"

func finsher() {
	fmt.Println("this is finish")

}
func largest(nums []int) {
	defer finsher()
	fmt.Println("start finding finish")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("largest number in ", nums, "is", max)

}

func main() {
	nums := []int{78, 45, 49, 235, 236}
	largest(nums)
}
