package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 基础
// 两数之和
// 考察：数组遍历、map使用
// 题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
// 链接：https://leetcode-cn.com/problems/two-sum/

func forArr(arr []int, target int) string {
	// 1.去除数组中比目标值大的数,并对数组根据数值大小排序,由低至高
to:
	for i, v := range arr {
		// 值大于target
		if v > target {
			arr = append(arr[:i], arr[i+1:]...)
			goto to
		}
	}
	sort.Ints(arr)
	fmt.Println("<<<<<<<<<<<< 小于目标值的数组:", arr)

	// 计算输出
	var str []string
	arrLen := len(arr) - 1
	for i, v := range arr {
		for j := range arrLen + 1 {
			if i >= j {
				continue
			}
			sum := v + arr[j]
			if sum == target {
				str = append(str, "[", strconv.Itoa(v), ",", strconv.Itoa(arr[j]), "]")
			}
			fmt.Println("外层:", v, "内层:", arr[j], "总合:", sum)
		}
	}
	// 输出数值转字符串
	return strings.Join(str, "")
}

func main66() {
	nums := []int{2, 7, 1, 11, 6, 5, 3, 10, 15}
	target := 9

	fmt.Printf("<<<<<<<<<<<<<< 最终两数相加满足目标值的集合为:%s,目标值:%s", forArr(nums, target), strconv.Itoa(target))
}
