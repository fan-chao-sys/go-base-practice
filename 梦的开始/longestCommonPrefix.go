package main

import "fmt"

// 最长公共前缀
// 考察：字符串处理、循环嵌套
// 题目：查找字符串数组中的最长公共前缀
// 链接：https://leetcode-cn.com/problems/longest-common-prefix/

func subStr(strArr []string) string {
	var publicStr string
	// 循环通过每个字符串前后对比出相同,依次对比,更新返回
	strArrLen := len(strArr)
	for i := range strArr {
		if i == strArrLen-1 {
			break
		}
		publicStr = publicStrPrefix(strArr[i], strArr[i+1])
	}
	return publicStr
}

// 切割返回公共字符串
func publicStrPrefix(preStr string, postStr string) string {
	preStrLen := len(preStr)
	postStrLen := len(postStr)

	var publicLen int
	if preStrLen > postStrLen {
		publicLen = postStrLen
	} else if postStrLen > preStrLen {
		publicLen = preStrLen
	}

	var num int
	for i := range publicLen {
		if preStr[i] != postStr[i] {
			break
		} else {
			num++
		}
	}
	return preStr[:num]
}

func main() {
	strOne := []string{"flower", "flow", "flight", "fl"} // 公共前缀: fl
	strTwo := []string{"dog", "racecar", "car"}          // 无
	fmt.Println("strOne公共字符为:", subStr(strOne))
	fmt.Println("strTWO公共字符为:", subStr(strTwo))
}
