package easycode

import (
	"math"
	"reflect"
)

// Policy 检查等差等比或者开方等差的策略
type Policy func(arr []int) (val int, ok bool)

// CalArray 数列之间的相加相减相乘相除
type CalArray func(arr []int) []int

// CalNum 此函数用于压栈保存状态
type CalNum func(a, b int) int

// Inference 所有策略保存在policy数组遍历检查/**
type Inference struct {
	policy []Policy
}

func NewInference(p ...Policy) *Inference {
	return &Inference{policy: p}
}

// Action 每次压栈之后需要保存的状态
type Action struct {
	method CalNum //method对应  add sub multiply quotient
	val    int    //val  作为第一个参数
}

//byZero 避免除0错误
func (action *Action) byZero() bool {
	sf1 := reflect.ValueOf(action.method)
	sf2 := reflect.ValueOf(quotient)
	return sf1.Pointer() == sf2.Pointer() && action.val == 0

}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func quotient(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func isInt(num float64) bool {
	return float64(int(num)) == num
}

//ArithmeticProgression 检查等差数列 通项公式：an = a1 + (n-1)d
func ArithmeticProgression(arr []int) (val int, ok bool) {
	n := len(arr)
	//fast path
	if n < 3 {
		return 0, false
	}
	for i := 1; i < n-1; i++ {
		//判断是否是等差数列
		if arr[i]*2 != arr[i-1]+arr[i+1] {
			return 0, false
		}
	}
	//公差
	difference := arr[n-1] - arr[n-2]

	return arr[0] + n*difference, true
}

//GeometricSequence 检查等比数列 an = a1 * q^n-1
func GeometricSequence(arr []int) (val int, ok bool) {
	n := len(arr)
	//fast path
	if n < 3 {
		return 0, false
	}
	for i := 1; i < n-1; i++ {
		//判断是否是符合要求的等比数列
		if arr[i] == 0 || arr[i-1] == 0 {
			return 0, false
		}
		if arr[i]*arr[i] != arr[i-1]*arr[i+1] || abs(arr[i]) < abs(arr[i-1]) {
			return 0, false
		}
	}
	q := arr[n-1] / arr[n-2]
	return arr[0] * int(math.Pow(float64(q), float64(n))), true
}

//SquareRootArithmeticSeries 检查开方等差数列
func SquareRootArithmeticSeries(arr []int) (val int, ok bool) {

	n := len(arr)
	//fast path
	if n < 3 {
		return 0, false
	}

	sqrt0 := math.Sqrt(float64(arr[0]))
	sqrt1 := math.Sqrt(float64(arr[1]))

	if !isInt(sqrt0) && !isInt(sqrt1) {
		return 0, false
	}
	difference := int(sqrt1) - int(sqrt0)
	ans := sqrt1
	for i := 2; i < n; i++ {
		if arr[i] < 0 {
			return 0, false
		}
		ans += float64(difference)
		if ans != math.Sqrt(float64(arr[i])) {
			return 0, false
		}
	}
	ans += float64(difference)
	return int(math.Pow(float64(ans), 2)), true
}

//differenceSequence 差分
func differenceSequence(arr []int) []int {
	n := len(arr)
	// res := make([]int, n-1)
	for i := 1; i < n; i++ {
		temp := arr[i-1]
		arr[i-1] = arr[i] - temp
	}
	return arr[0 : n-1]
}

//quotientSequence 商
func quotientSequence(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		if arr[i-1] == 0 || arr[i]%arr[i-1] != 0 {
			return nil
		}
		temp := arr[i-1]
		arr[i-1] = arr[i] / temp
	}
	return arr[0 : n-1]
}

//addSequence  相加
func addSequence(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		arr[i-1] = arr[i] + arr[i-1]
	}
	return arr[0 : n-1]
}

//multiplySequence 相乘
func multiplySequence(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		temp := arr[i-1]
		if temp == 0 {
			return nil
		}
		arr[i-1] = arr[i] * temp
	}
	return arr[0 : n-1]
}

//complex 组合的策略
func complex(arr []int, stack *Stack, inf *Inference) (int, bool) {
	//观察数列数字间的变化幅度的大小，如果前几项较小，末项却突然增大数倍，可以考虑等比数列；
	//如果数列的起伏不大，变化幅度小且逐渐递增或递减，则可考虑等差数列。
	if len(arr) <= 3 {
		return 0, false
	}
	n := len(arr)
	begin, end := n-2, n-1

	//如题：优先是差分 商分
	//如果数字间浮动很大，考虑取商
	if arr[begin] != 0 && arr[end]/arr[begin] > 4 {
		stack.Push(&Action{multiply, arr[n-1]})
		arr = quotientSequence(arr)
	} else {
		//反之取差
		stack.Push(&Action{add, arr[n-1]})
		arr = differenceSequence(arr)
	}
	if arr == nil {
		return 0, false
	}

	//计算
	if val, ok := cal(arr, stack, inf); ok {
		return val, ok
	}

	//数组长度 > 3 还有分析的空间则递归执行
	return complex(arr, stack, inf)
}

//complexSingleCondition 对应其他需要单个检查的规则，额外添加的
func complexSingleCondition(arr []int, stack *Stack, calArray CalArray, calNum CalNum, inf *Inference) (int, bool) {
	n := len(arr)
	if len(arr) <= 3 {
		return 0, false
	}
	stack.Push(&Action{calNum, arr[n-1]})
	arr = calArray(arr)

	if arr == nil {
		return 0, false
	}

	if val, ok := cal(arr, stack, inf); ok {
		return val, ok
	}

	return complexSingleCondition(arr, stack, calArray, calNum, inf)
}

//cal 从栈中顺序pop 推算出下一项值
func cal(arr []int, stack *Stack, inf *Inference) (int, bool) {
	if val, ok := check(arr, inf); ok {
		action := stack.Pop()
		ans := action.method(val, action.val)
		for !stack.IsEmpty() {
			action = stack.Pop()
			//禁止除0
			if action.byZero() {
				return 0, false
			}
			ans = action.method(ans, action.val)
		}
		return ans, ok
	}

	return 0, false
}

//check 简单循环检查
func check(arr []int, inf *Inference) (int, bool) {
	for _, f := range inf.policy {
		if val, ok := f(arr); ok {
			return val, ok
		}
	}
	return 0, false

}

func checkList(list []int) (val int, ok bool) {
	var _list []int = make([]int, len(list))
	copy(_list, list)
	//fast path 简单数列不需要变化
	inf := NewInference(ArithmeticProgression, GeometricSequence, SquareRootArithmeticSeries)
	if val, ok = check(list, inf); ok {
		return val, ok
	}

	//复杂数列 递归分支判断
	complexStack := NewStack()
	if val, ok = complex(_list, complexStack, inf); ok {
		return val, ok
	}

	addStack := NewStack()
	var addList []int = make([]int, len(list))
	copy(addList, list)
	//add分支
	if val, ok = complexSingleCondition(addList, addStack, addSequence, sub, inf); ok {
		return val, ok
	}

	//mul分支
	mulStack := NewStack()
	var mulList []int = make([]int, len(list))
	copy(mulList, list)
	if val, ok = complexSingleCondition(mulList, mulStack, multiplySequence, quotient, inf); ok {
		return val, ok
	}

	return 0, false

}

//Stack 代码需要用到的辅助栈，计算递推成规律数列之后反向逆推
type Stack struct {
	actions []*Action
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.actions) == 0
}

func (stack *Stack) Push(action *Action) {
	stack.actions = append(stack.actions, action)
}

func (stack *Stack) Pop() *Action {
	if !stack.IsEmpty() {
		res := stack.actions[len(stack.actions)-1]
		stack.actions = stack.actions[0 : len(stack.actions)-1]
		return res
	}

	return nil
}
