package main

import "sort"

// 825. 适龄的朋友 https://leetcode-cn.com/problems/friends-of-appropriate-ages/
// 难度：中等
func numFriendRequests(ages []int) (ans int) {
	sort.Ints(ages)
	n := len(ages)
	for k, i, j := 0, 0, 0; k < n; k++ {
		//不满足的情况
		for i < k && !check(ages[i], ages[k]) {
			i++
		}
		if j < k {
			j = k
		}
		for j < n && check(ages[j], ages[k]) {
			j++
		}
		if j > i {
			ans += j - i - 1
		}
	}

	return
}

func check(x, y int) bool {
	if y <= x/2+7 {
		return false
	}
	if y > x {
		return false
	}
	if y > 100 && x < 100 {
		return false
	}
	return true
}
