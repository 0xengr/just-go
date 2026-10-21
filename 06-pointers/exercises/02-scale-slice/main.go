package main

import "fmt"

func scaleFirst(nums []int) {
	nums[0] = nums[0] * 10
}

func main() {
	nums := []int{1, 2, 3}
	scaleFirst(nums)
	fmt.Println(nums)
}
