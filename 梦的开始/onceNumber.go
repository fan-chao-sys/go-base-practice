package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 控制流程
// 只出现一次的数字：
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。

func forArrMap(arr []int) string {
	arrMap := make(map[int]int)
	for i := range arr {
		if _, ok := arrMap[arr[i]]; ok {
			arrMap[arr[i]]++
		} else {
			arrMap[arr[i]] = 1
		}
	}
	fmt.Println(arrMap)

	// 取1次字符串
	var str []string
	for k, v := range arrMap {
		if v == 1 {
			str = append(str, strconv.Itoa(k))
		}
	}
	return strings.Join(str, ",")
}

func main777() {
	onceArr := []int{2, 7, 7, 11, 11, 5, 5, 10, 10}
	sort.Ints(onceArr)
	fmt.Printf("出现1次的整数值为:%s\n", forArrMap(onceArr))
}
