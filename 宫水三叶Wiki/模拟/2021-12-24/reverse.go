package main

import "math"

//https://leetcode-cn.com/problems/reverse-integer/
func reverse(x int) (rev int) {
	//fast path
	//主要是解决溢出问题
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}

	return rev
}
