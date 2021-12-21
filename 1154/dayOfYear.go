package main

import "strconv"

//其实写这么多代码主要就是把一些方法抽出去了
func dayOfYear(date string) int {
	year := getYear(date)
	month := getMonth(date)
	day := getDay(date)

	ret := 0
	for i := 1; i < month; i++ {
		ret += getDayByMonth(i, year)
	}

	return ret + day
}

//通过月份计算当前日期
func getDayByMonth(month, year int) int {

	if month == 2 {
		return m29(year)
	}

	if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
		return 31
	}

	return 30
}

//闰年平年
func m29(year int) int {
	if year%4 == 0 {
		return 29
	}

	return 28
}

//获取年份
func getYear(date string) int {
	year := date[0:4]
	val, err := strconv.Atoi(year)
	if err != nil {
		return 0
	}
	return val
}

//获取月份
func getMonth(date string) int {
	month := date[5:7]
	val, err := strconv.Atoi(month)
	if err != nil {
		return 0
	}
	return val
}

//获取日期
func getDay(date string) int {
	day := date[8:10]
	val, err := strconv.Atoi(day)
	if err != nil {
		return 0
	}
	return val
}
