// Package Stack
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/1 14:36
// -----------------------------------------------
package Stack

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := make([]int, 0)

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) != 0 {
			res[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return res
}
