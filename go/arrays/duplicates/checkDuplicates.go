package main

func containsDuplicate(nums []int) bool {
	nonDup := make(map[int]bool)
	for _, val := range nums {
		if !nonDup[val] {
			nonDup[val] = true
		} else {
			return true
		}
	}
	return false
}
