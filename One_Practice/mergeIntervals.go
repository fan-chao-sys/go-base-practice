package main

import (
	"fmt"
	"sort"
)

// 合并区间
// 数组 intervals 表示若干个区间集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请合并所有重叠区间，返回一个不重叠区间数组，该数组需恰好覆盖输入中的所有区间。
// 可先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
// 将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。

func merge(intervals [][]int) [][]int {
	// 1. 按区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 2. 创建结果切片
	var result [][]int

	// 3. 遍历排序后的区间
	for _, interval := range intervals {
		// 如果结果为空，或者当前区间与最后一个区间不重叠
		if len(result) == 0 || interval[0] > result[len(result)-1][1] {
			result = append(result, interval)
		} else {
			// 有重叠，合并区间（更新最后一个区间的结束位置）
			if interval[1] > result[len(result)-1][1] {
				result[len(result)-1][1] = interval[1]
			}
		}
	}

	return result
}

func main88() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Print(merge(intervals))
}
