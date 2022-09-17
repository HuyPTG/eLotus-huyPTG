package main

//Maximum Length of Repeated Subarray
import "fmt"

func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	maxLen := 0

	for i := len(nums1) - 1; i >= 0; i-- {
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
				}
			}
		}
	}

	return maxLen
}

func main() {
	var n1 []int
	fmt.Scan(&n1)
	var n2 []int
	fmt.Scan(&n2)
	max := findLength(n1, n2)
	fmt.Println("Find Length:", max)
}
