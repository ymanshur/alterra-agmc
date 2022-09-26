package main

import (
	"fmt"
	"sort"
)

func main() {
	// var val = []int32{5, -1, -2, -3, 8, 7}
	// fmt.Println(findLongestSubsequence(val))
	// fmt.Println(getMaximumEvenSum(val))

	var a = []int32{1, 2, 3, 4, 5}
	var b = []int32{1, 1, 1, 6, 6}
	fmt.Println(findNumOfPairs(a, b))
}

/**
 * Event Difference - HackerRank
 *
 * Hint:
 * To have the sum of adjacent elements even, the difference between the first and last elements has to be even in the sorted array.
 *      1. Sort the array.
 *      2. Find the first and last occurrence of the even number in the array.
 *      3. Find the first and last occurrence of the odd number in the array.
 * The answer is the maximum difference between the first and last occurrence of the even number in the array and the first and last occurrence of the odd number in the array.
 */

func findLongestSubsequence(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	arrLength := int32(len(arr))
	var firstOddIndex, lastOddIndex, firstEvenIndex, lastEvenIndex, i int32
	for i = 0; i < arrLength; i++ {
		if arr[i]%2 == 0 && lastEvenIndex != i {
			lastEvenIndex = i
		} else if arr[i]%2 != 0 {
			lastOddIndex = i
		}
	}

	for i := arrLength - 1; i >= 0; i-- {
		if arr[i]%2 == 0 {
			firstEvenIndex = i
		} else if arr[i]%2 != 0 {
			firstOddIndex = i
		}
	}

	if (lastEvenIndex - firstEvenIndex) > (lastOddIndex - firstOddIndex) {
		return int32(lastEvenIndex - firstEvenIndex + 1)
	} else {
		return int32(lastOddIndex - firstOddIndex + 1)
	}
}

/**
 * Discount Tags - HackerRank
 *
 * Hint:
 *      1. First, it will always be optimal to choose as many positive elements as possible.
 *      2. Then, to make the sum even, there are two options, either remove an odd positive element or add a negative odd element.
 *      3. So, find the minimum amongst the absolute values of all odd elements, and update the answer accordingly.
 */

func getMaximumEvenSum(val []int32) int64 {
	var maxSum int64
	var minOddNumber int32

	for _, v := range val {
		if v > 0 {
			maxSum += int64(v)
		}
		if v%2 != 0 {
			if minOddNumber == 0 {
				minOddNumber = v
			} else if Abs(v) < Abs(minOddNumber) {
				minOddNumber = v
			}
		}
	}

	if maxSum%2 != 0 {
		if minOddNumber < 0 {
			maxSum += int64(minOddNumber)
		} else {
			maxSum -= int64(minOddNumber)
		}
	}

	return maxSum
}

func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func findNumOfPairs(a []int32, b []int32) int32 {
	var maxPairNum int32
	// var startBIndex int
	var medianBIndex int
	var n = len(a)

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })

	fmt.Println(a)
	fmt.Println(b)

	for i := 0; i < n; i++ {
		low := 0
		high := n - 1

		for 0 <= high {
			medianBIndex = (low + high) / 2

			if b[medianBIndex] < a[i] {
				maxPairNum += 1
				break
			} else {
				high = medianBIndex - 1
			}
		}
		// for j := startBIndex; j < n; j++ {
		// 	if a[i] > b[j] {
		// 		maxPairNum += 1
		// 		startBIndex = j + 1
		// 		break
		// 	}
		// }
	}

	return maxPairNum
}
