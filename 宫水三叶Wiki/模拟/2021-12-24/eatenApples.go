package main

import "container/heap"

// https://leetcode-cn.com/problems/maximum-number-of-eaten-apples/
//用golang的堆是真的挺费劲的 我实话说
func eatenApples(apples []int, days []int) int {
	var h hp
	var ans int
	var time int
	for time = 0; time < len(apples); time++ {
		for len(h) > 0 && h[0].end <= time {
			heap.Pop(&h)
		}
		if apples[time] > 0 {
			heap.Push(&h, pair{time + days[time], apples[time]})
		}
		if len(h) > 0 {
			h[0].left--
			if h[0].left == 0 {
				heap.Pop(&h)
			}
			ans++
		}
	}

	for len(h) > 0 {
		for len(h) > 0 && h[0].end <= time {
			heap.Pop(&h)
		}

		if len(h) == 0 {
			break
		}

		p := heap.Pop(&h).(pair)
		//最多吃多少天 和一共多少个 ，最多只能吃这么多 不能再多了
		num := min(p.end-time, p.left)
		ans += num
		time += num

	}

	return ans
}

//这里是直接拷贝官方的 懒得写了
type pair struct{ end, left int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
