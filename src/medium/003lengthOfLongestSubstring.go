package main

import "fmt"

/*3  给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
示例1 ： 输入: "abcabcbb"   输出: 3  ；  解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例2 ：输入: "pwwkew"    输出: 3  ；
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
func lengthOfLongestSubstring(s string) int {
	var k, max int
	j, i := 0, 0
	length := len(s)
	for ; j < length; j++ {
		for k = i; k < j; k++ {
			if s[k] == s[j] {
				i = k + 1
				break
			}
		}
		if j-i+1 > max {
			max = j - i + 1
		}
	}
	return max
}

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}