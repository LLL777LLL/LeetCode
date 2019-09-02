package main

import "fmt"

/*给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
示例1： 输入: "babad"  输出: "bab"  ;
注意: "aba" 也是一个有效答案。

示例2： 输入: "cbbd"   输出: "bb"   ;
*/

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		l := 0
		if len1 > len2 {
			l = len1
		} else {
			l = len2
		}
		if l > end-start {
			start = i - (l-1)/2
			end = i + l/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	l, r := left, right
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

func main() {
	str := "cdcbaabcd"
	fmt.Println(longestPalindrome(str))
}