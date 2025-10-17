package main

import (
	"fmt"
	"strconv"
)

// 基本值类型 -加一
// 考察：数组操作、进位处理
// 题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
// 链接：https://leetcode-cn.com/problems/plus-one/

// 举例: 数组 [1, 2, 3]，表示大整数 123，加 1 后变成 124，返回数组是 [1, 2, 4]；
//		再比如数组 [9, 9]，表示大整数 99，加 1 后变成 100，返回数组是 [1, 0, 0]。

func calculate(arr []int) []int {

	// 取数组所有数值,转为int 并+1,后转换回数组
	var result int
	var arrLen = len(arr)
	zero := zeroNum(arrLen)

	var sign int = 1
	for i := range arr {
		if sign == 1 {
			result += arr[i] * zero
		} else if sign == arrLen {
			result += arr[i]
		} else if sign >= 2 {
			var zeros int
			for j := 0; j < sign; j++ {
				zeros = zero / 10
			}
			result += arr[i] * zeros
		}
		sign++
	}

	result += 1
	fmt.Println("num:计算结果:", result)

	// 将整数转换为字符串
	str := strconv.Itoa(result)
	// 创建结果数组
	result2 := make([]int, 0, len(str))
	// 遍历字符串，将每个字符转换为int
	for _, c := range str {
		num, _ := strconv.Atoi(string(c))
		result2 = append(result2, num)
	}

	return result2
}

func zeroNum(len int) int {
	var zero int = 1 // 获取百位数
	if len == 2 {
		zero = 10
	} else {
		for i := 0; i < len-1; i++ {
			zero *= 10
		}
	}
	return zero
}

func main() {
	var arr = [3]int{1, 2, 3}
	var arr2 = [2]int{9, 9}

	fmt.Println(calculate(arr[:]))
	fmt.Println(calculate(arr2[:]))
}
