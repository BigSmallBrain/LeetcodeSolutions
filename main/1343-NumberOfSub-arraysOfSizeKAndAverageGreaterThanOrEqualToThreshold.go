// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/12 14:55
// -----------------------------------------------
package main

func numOfSubarrays(arr []int, k int, threshold int) int {
	n := len(arr)
	sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}

	res := 0
	kThresholds := k * threshold
	if sum >= kThresholds {
		res++
	}
	for index := k; index < n; index++ {
		sum = sum + arr[index] - arr[index-k]
		if sum >= kThresholds {
			res++
		}
	}
	return res
}
