package main

func repeatedStringMatch(a string, b string) int {
	var s string = a
	var count int = 1
	var bl int = len(b)
	for len(s) < bl {
		s += a
		count++
	}

	var snew int = len(s) + len(a)
	var sold int = len(s)

	//按照分析那么我提前使用 snew 表示 c1 + 1的长度
	for i := 0; i < snew; i++ {
		//b的长度加上 i 不能大于 snew的长度
		if bl+i < snew && s[i:bl+i] == b { //如果发现等于的子串 则直接返回c1
			return count
		}

		//如果 字符串 b的长度 加上 i 长度等于 sold的长度
		//那么意味着 c1不满足 需要扩容成c2 此时 count+1
		if bl+i == sold {
			s += a
			count++
		}
	}

	return -1
}
