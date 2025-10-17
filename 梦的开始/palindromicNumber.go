package main

import (
	"fmt"
	"strconv"
)

// 回文数
// 考察：数字操作、条件判断
// 题目：判断一个整数是否是回文数

func palindromic(num int) bool {
	// 负数 和 单位数 都是回文数
	if num >= -9 && num <= 9 {
		return true
	}

	// 整数转字符串
	var strNum = strconv.Itoa(num)
	var lenNum = len(strNum)
	var middle = len(strNum) / 2
	// 取模判断奇偶数
	if lenNum%2 == 1 { // 奇数
		middles := middle + 1 //  中间数，如：123的2, 3434341的4
		fmt.Println("奇数还没写", middles)
	} else { // 偶数
		var num int
		var frequency int
		for i := range middle {
			// 根据中间位置对比前后简介的字符是否相等？
			num++
			if strNum[i] == strNum[lenNum-num] {
				frequency++
			}
		}
		if frequency == middle {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("1:", palindromic(1))
	fmt.Println("2:", palindromic(121))
	fmt.Println("3:", palindromic(12445))
	fmt.Println("4:", palindromic(3223))
	fmt.Println("5:", palindromic(1564224651))
}
