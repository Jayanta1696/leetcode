package main

func twoSumHash(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if _, ok := hashmap[complement]; ok {
			return []int{hashmap[complement], i}
		}
		hashmap[nums[i]] = i
	}
	return []int{}



}	