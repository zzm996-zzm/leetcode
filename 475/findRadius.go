package main

import "sort"

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	var r int = 0
	var n int = len(heaters)
	var cur int = 0

	for i := 0; i < len(houses); i++ {

		for cur < n {
			//找到当前房子距离下一个暖气的距离
			if heaters[cur] >= houses[i] {
				break
			}
			cur++
		}
		//如果当前房子是处于中间位置，那么计算 上一个暖气和下一个暖气之间哪个最小
		//再判断是否比之前的大，因为要保证所有房子都能覆盖到暖气
		if cur > 0 && cur < n {
			r = max(r, min(heaters[cur]-houses[i], houses[i]-heaters[cur-1]))
		}

		//边界处理 避免 cur-1 = -1
		if cur == 0 {
			r = max(r, heaters[cur]-houses[i])
		}

		//边界处理 避免 cur == n
		if cur == n {
			r = max(r, houses[i]-heaters[cur-1])
		}

	}

	return r

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
