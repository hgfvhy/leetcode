package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 17
	fmt.Print(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		p, ok := hashTable[target-x]
		if ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func twoSumMine(nums []int, target int) []int {
	numsLen := len(nums)
	res := make([]int, 0)

fore:
	for i := 0; i <= numsLen-2; i++ {
		for j := i + 1; j <= numsLen-1; j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)
				break fore
			}
		}
	}
	return res
}
