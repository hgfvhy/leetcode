package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "eqeaaqbcd"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring01(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	//fmt.Println(m)
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		fmt.Println(m)
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			fmt.Println(m, s[rk+1])
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
