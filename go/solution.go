package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) (int, int) {
	sort.Ints(nums)
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]+nums[j] == target {
			return i, j
		} else if nums[i]+nums[j] < target {
			i++
		} else {
			j--
		}
	}
	return 0, len(nums) - 1
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	var res [][]int = make([][]int, 100)
	for i := 0; i < len(nums); i++ {
		if i >= 1 && nums[i-1] == nums[i] {
			continue
		}
		for p, q := i+1, len(nums)-1; p < q; {
			if nums[p]+nums[q]+nums[i] == 0 {
				res = append(res, []int{i, p, q})
				p++
				q--
			} else if nums[p]+nums[q]+nums[i] < 0 {
				p++
			} else {
				q--
			}
		}
	}
	return res
}

func main() {
	fmt.Println("hello")
	//	var nums = []int{-1, 0, 1, 2, -1, -4}
	//	fmt.Println(len(nums))
	//threeSum([]int{-1, 0, 1, 2, -1, -4})
}
