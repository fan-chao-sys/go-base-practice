package main

import (
	"fmt"
)

// 字符串 - 有效的括号
// 考察：字符串处理、栈的使用
// 题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
// 要求：
// 		有效字符串需满足：
//		左括号必须用相同类型的右括号闭合。
//		左括号必须以正确的顺序闭合。
//		每个右括号都有一个对应的相同类型的左括号。
// 链接：https://leetcode-cn.com/problems/valid-parentheses/

// 思路：
// 1.由于索引下标从0开始,所以字符串长度一定是单数。
// 2.已知固定字符有: '('  ')' '['  ']', 根据字符串的长度,循环判断上一个字符串是否包括固定结尾括号即可。 最终都满足输出true

// 校验字符串长度
func isValid(s string) bool {
	fmt.Println("--------------------- 开始 --------------------")

	// 获取字符串长度 除2 取模? 0 = 偶数 / 1 = 奇数
	var length = len(s)
	fmt.Println("字符串:", s, ",长度为:", length)
	if length != 0 {
		if length%2 == 0 {
			fmt.Println("当前字符串:", s, "是偶数!")

			var pair = length / 2 // 总共几对值
			for i := range s {
				fmt.Printf("循环数组元素:%c\n", s[i])
				if s[i] == '(' && s[i+1] == ')' {
					pair = pair - 1
				}
				if s[i] == '[' && s[i+1] == ']' {
					pair = pair - 1
				}
			}

			if pair <= 0 {
				return true
			}
		}
	} else {
		fmt.Print("字符串为空!")
	}
	return false
}

func main() {
	fmt.Println("第1个:", isValid("()"))
	fmt.Println("第2个:", isValid(""))
	fmt.Println("第3个:", isValid("()["))
	fmt.Println("第4个:", isValid("()]"))
	fmt.Println("第5个:", isValid("()[]"))
	fmt.Println("第6个:", isValid("(]"))
	fmt.Println("第7个:", isValid("()[[[]"))
	fmt.Println("第8个:", isValid("()[()][]"))
}

// ------------------------------------------------------------------  实现方式2： 栈 + 哈希表（推荐，易维护）
// 核心原理：利用栈的「后进先出」特性匹配括号 —— 遇到左括号入栈，遇到右括号时检查栈顶是否为对应的左括号，匹配则出栈，不匹配则无效。
// 哈希表作用：提前存储「右括号 - 左括号」的对应关系，避免复杂的条件判断，提升代码简洁性和可维护性。
// 边界处理：
// 若字符串长度为奇数，直接返回 false（无法成对匹配）。
// 遍历结束后，栈需为空（所有左括号均找到匹配的右括号）。

// func isValid(s string) bool {
//    // 1. 边界优化：奇数长度的字符串必然无法匹配，直接返回false
//    if len(s) % 2 != 0 {
//        return false
//    }
//
//    // 2. 定义右括号到左括号的映射表，简化匹配判断
//    bracketMap := map[byte]byte{
//        ')': '(', // 右括号')'对应左括号'('
//        ']': '[', // 右括号']'对应左括号'['
//        '}': '{', // 右括号'}'对应左括号'{'
//    }
//
//    // 3. 初始化栈（用切片模拟栈，append入栈，len(stack)-1取栈顶，stack[:len(stack)-1]出栈）
//    stack := make([]byte, 0)
//
//    // 4. 遍历字符串中的每个字符
//    for i := 0; i < len(s); i++ {
//        char := s[i]
//        // 5. 判断当前字符是否为右括号（存在于映射表的key中）
//        if matchLeft, isRightBracket := bracketMap[char]; isRightBracket {
//            // 5.1 若栈为空，或栈顶元素不是对应的左括号，直接返回false
//            if len(stack) == 0 || stack[len(stack)-1] != matchLeft {
//                return false
//            }
//            // 5.2 匹配成功，弹出栈顶元素（左括号）
//            stack = stack[:len(stack)-1]
//        } else {
//            // 6. 若为左括号，直接入栈
//            stack = append(stack, char)
//        }
//    }
//
//    // 7. 遍历结束后，栈为空说明所有括号都匹配成功；否则存在未匹配的左括号
//    return len(stack) == 0
//}

// ------------------------------------------------------------------------------   实现方式3：纯栈 + 条件判断（无哈希表，基础实现）
// 核心原理：与解法一一致，均基于栈的「后进先出」特性。
// 差异点：不使用哈希表，而是通过 if-else 条件判断右括号对应的左括号，适合理解基础逻辑。
// 适用场景：括号类型较少（本题仅 3 种）的场景，代码直观但扩展性稍弱（若增加括号类型需新增条件）。

// func isValid(s string) bool {
//    // 1. 边界优化：奇数长度直接返回false
//    if len(s) % 2 != 0 {
//        return false
//    }
//
//    // 2. 用切片模拟栈，存储左括号
//    stack := make([]byte, 0)
//
//    // 3. 遍历字符串每个字符
//    for i := range s {
//        char := s[i]
//        // 4. 左括号：直接入栈
//        if char == '(' || char == '[' || char == '{' {
//            stack = append(stack, char)
//        } else {
//            // 5. 右括号：先检查栈是否为空（无匹配的左括号）
//            if len(stack) == 0 {
//                return false
//            }
//            // 6. 弹出栈顶元素，判断是否与当前右括号匹配
//            top := stack[len(stack)-1]
//            stack = stack[:len(stack)-1] // 出栈
//            // 7. 逐一判断右括号对应的左括号是否匹配
//            if (char == ')' && top != '(') ||
//               (char == ']' && top != '[') ||
//               (char == '}' && top != '{') {
//                return false
//            }
//        }
//    }
//
//    // 8. 栈为空则所有括号匹配成功
//    return len(stack) == 0
//}
