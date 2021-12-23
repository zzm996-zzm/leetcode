package main

type substrHash func(left, right int, s string) int

type check func(s string, len int) string

const power = 31

func longestDupSubstring(s string) string {
	var n int = len(s)
	var p, h []int

	p = make([]int, n+1)
	h = make([]int, n+1)

	p[0] = 1

	for i := 1; i <= n; i++ {
		h[i] = h[i-1]*power + int(s[i-1])
		p[i] = p[i-1] * power
	}

	substrHash := func(left, right int, h, p []int) int {
		return h[right] - h[left]*p[right-left]
	}

	check := func(s string, tl int) string {
		var n int = len(s)
		m := make(map[int]int)
		for i := 0; i+tl < n; i++ {
			right := tl + i + 1
			hash := substrHash(i, right, h, p)
			if _, has := m[hash]; has {
				return s[i:right]
			}
			m[hash]++
		}

		return ""
	}

	//二分结果

	l, r := 0, len(s)
	var str string
	var res string
	for l < r {
		mid := (l + r) >> 1
		str = check(s, mid)
		if str == "" {
			r = mid
		} else {
			l = mid + 1
		}

		if len(str) > len(res) {
			res = str
		}
	}

	return res

}
