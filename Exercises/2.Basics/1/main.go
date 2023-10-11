package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(doubleDetector(nums))
}

func doubleDetector(nums []int) bool {
	for _, v := range nums {
		var count int
		for _, n := range nums {
			if v == n {
				count++
			}
			if count >= 2 {
				return true
			}
		}
	}
	return false
}
