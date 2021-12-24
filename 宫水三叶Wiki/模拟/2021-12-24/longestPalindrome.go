package main

//解法1 暴力解法
// https://leetcode-cn.com/problems/longest-palindromic-substring/submissions/
func longestPalindrome(s string) string {
	var l, r int
	var ans string
	for i := 0; i < len(s); i++ {
		//奇数
		l, r = i-1, i+1
		ans = max(ans, getString(s, l, r))
		l, r = i-1, i
		ans = max(ans, getString(s, l, r))
	}

	return ans
}

func max(sub, ans string) string {
	if len(sub) >= len(ans) {
		return sub
	}

	return ans
}

func getString(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

//解法2 Manacher

func longestPalindrome2(s string) string {
	//fast path
	if len(s) == 1 {
		return s
	}
	chars := manacherString(s)
	n := len(chars)
	pArr := [n]int{}

	C, R, pos := -1, -1, -1
	max := -1 << 31

	for i := 0; i < n; i++ {
		if i < R {
			pArr[i] = min(pArr[C*2-i], R-i)
		} else {
			pArr[i] = 1
		}

		for i+pArr[i] < n && i-pArr[i] > -1 {
			if chars[i+parr[i]] == chars[i-pArr[i]] {
				pArr[i]++
			} else {
				break
			}
		}

		if i+pArr[i] > R {
			R = i + pArr[i]
			C = i
		}

		if pArr[i] > max {
			max = pArr[i]
			pos = i
		}
	}

	offset := pArr[pos]
	ans := make([]byte, 0)
	for i := pos - offset + 1; i <= pos+offset-1; i++ {
		if chars[i] != '#' {
			ans = append(ans, chars[i])
		}
	}

	return string(ans)
}

func manacherString(s string) []byte {
	chars := make([]byte, len(s)*2+1)
	for i, index := 0, 0; i < len(chars); i++ {
		//字符串转换，添加占位符
		if (i & 1) == 0 {
			char[i] = '#'
		} else {
			char[i] = s[index]
			index++
		}
	}

	return chars
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
