package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {

	dp := [2][]int{} //array has to always be initilized
	dp[1] = make([]int, len(nums))
	dp[0] = make([]int, len(nums))

	dp[1][0] = nums[0]
	dp[0][0] = dp[1][0]

	for i := 1; i < len(nums); i++ {
		fmt.Println(nums[i], nums[i]+dp[1][i-1])
		dp[1][i] = int(math.Max(float64(nums[i]), float64(nums[i]+dp[1][i-1])))
		dp[0][i] = int(math.Max(float64(dp[0][i-1]), float64(dp[1][i])))
	}
	return dp[0][len(dp[0])-1]
}
