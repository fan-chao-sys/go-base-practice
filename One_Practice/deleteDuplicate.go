package main

import (
	"fmt"
	"sort"
)

// 引用类型: 切片
// 删除有序数组中的重复项：有序数组 nums ，删除重复元素，使每个元素只出现一次，返回删除后元素个数综合及新数组排序后的长度。
//					  不要使用额外数组空间，必须原修改输入数组并使用 O(1) 额外空间条件下完成。
//					  可以使用双指针法：
//						  一个慢指针 i 用于记录不重复元素的位置，
//						  一个快指针 j 用于遍历数组，
//						  当 nums[i] 与 nums[j] 不相等时，
//						  将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
// 链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

// 一个有序数组,删除重复元素并不可影响原有排序,最后输出元素个数及新数组。
func delRepeat(arr []int) {
	arrMap := make(map[int]bool)
	var result []int
	for i := range arr {
		if !arrMap[arr[i]] {
			// 如果不存在，添加到map和结果中
			arrMap[arr[i]] = true
			result = append(result, arr[i])
		} else {
			// 如果已存在，false（相当于删除重复项）
			arrMap[arr[i]] = false
		}
	}
	fmt.Println("result:", result)

	var a int
	for i := range arrMap {
		if !arrMap[i] {
			a++
		}
	}
	fmt.Println("arrMa重复个数:", a)
	fmt.Println("arrMa重复map:", arrMap)
}

func main10() {
	var arr = []int{1, 2, 2, 3, 4, 4, 5, 6, 6, 7, 8, 8}
	sort.Ints(arr)
	delRepeat(arr)
}
