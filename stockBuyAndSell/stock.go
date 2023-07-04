package main

import (
	"math"
)

/*
	type day struct {
	    minDay int
	    maxday int
	    profit int
	}
*/
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	min := prices[0]
	max := 0
	//	var days []day
	//	minDay := 1
	//	maxDay := 0
	maxProfit := math.MinInt
	for _, val := range prices[1:] {
		if val < min {
			//	days = append(days, day{minDay, maxDay, max - min})
			if maxProfit < max-min {
				maxProfit = max - min
			}
			min = val
			max = 0
			//			minDay = i + 1
			continue
		}

		if max < val {
			max = val
			//			maxDay = i + 1
			if maxProfit < max-min {
				maxProfit = max - min
			}
		}
	}
	/*if maxDay != 0 {
		days = append(days, day{minDay, maxDay, max - min})
	}*/

	//fmt.Println(days)
	if maxProfit > 0 {
		return maxProfit
	}
	return 0
}
