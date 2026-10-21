package main

import "fmt"

func main() {
	nums := []int{3, 1, 2}
	nums = append(nums, 4)
	dst := make([]int, len(nums))
	copy(dst, nums)
	clear(dst)
	fmt.Println("original:", nums)
	fmt.Println("cleared copy:", dst)
}
