package main

import (
	"fmt"
)

func main() {
	var val = []int32{5, 87, 99, 85, 50, 93}
	fmt.Println(findLongestSubsequence(val))
	// fmt.Println(getMaximumEvenSum(val))
}

func findLongestSubsequence(arr []int32) int32 {
	longestSubset := []int32{}

	length := len(arr)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []int32
		for object := 0; object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, arr[object])
			}
		}

		subsetLength := len(subset)
		longestSubsetLength := len(longestSubset)
		min, max := findMinMax(subset)
		sumOfDiff := max - min
		if subsetLength > longestSubsetLength && sumOfDiff%2 == 0 {
			fmt.Println(subset)
			longestSubset = subset
		}
	}

	fmt.Println(longestSubset)
	return int32(len(longestSubset))
}

func findMinMax(arr []int32) (int32, int32) {
	min := arr[0]
	max := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func getMaximumEvenSum(val []int32) int64 {
	var maxSumOfVal int32
	maxSumOfVal = 0

	length := len(val)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []int32
		for object := 0; object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, val[object])
			}
		}

		sumOfVal := sum(subset)
		if sumOfVal%2 == 0 && sumOfVal > maxSumOfVal {
			maxSumOfVal = sumOfVal
		}
	}

	return int64(maxSumOfVal)
}

func sum(val []int32) int32 {
	var total int32
	total = 0
	for _, v := range val {
		total += v
	}
	return total
}
