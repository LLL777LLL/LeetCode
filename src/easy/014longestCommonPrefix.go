package main

import (
	"fmt"
	"time"
)

/*编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
示例1:  输入: ["flower","flow","flight"]    输出: "fl" ;
示例2:  输入: ["dog","racecar","car"]       输出: ""   ;   解释: 输入不存在公共前缀。
说明:   所有输入只包含小写字母 a-z 。
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	str := strs[0]
	for i := 0; i < len(strs)-1; i++ {
		j, k := 0, 0
		for k < len(str) && k < len(strs[i+1]) {
			k++ //判断循环次数
			if str[:j+1] == strs[i+1][:j+1] {
				j++ //记录相同前缀字符的个数
			}
		}
		if j < 1 {
			return ""
		}
		str = str[:j] //暂存相同前缀，用于下次比较
	}
	return str
}

//这个方法在goLand运行正确，但是到 LeetCode上就报错：超出时长，看了官方的思路后有了以上的解法（建议放弃这种方式）
func longestCommonPrefix2(strs []string) string {
	t := time.Now()
	i := 0
	isBreak := false
	for {
		l := 0
		for _, v := range strs {
			if i == len(v) {
				isBreak = true
				break
			} else if strs[0][:i+1] == v[:i+1] {
				l++
			} else {
				isBreak = true
				break
			}
		}
		if l == len(strs) {
			i++
		}
		if isBreak {
			break
		}
	}
	fmt.Println(time.Since(t))
	if i > 0 {
		return strs[0][:i]
	} else {
		return ""
	}
}

func main() {
	str := []string{
		//"dog","racecar","car",
		"flower", "flow", "flowight",
	}
	fmt.Println(longestCommonPrefix(str))
}
