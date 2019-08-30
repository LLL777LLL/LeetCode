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
	var str string
	var l int
	for i:=0;i < len(strs);i++ {
		
	}
}



//这个方法在goLand运行正确，但是到 LeetCode上就报错：超出时长，看了官方的思路后有了以上的解法
func longestCommonPrefix2(strs []string) string {
	t := time.Now()
	i := 0
	isBreak := false
	for {
		l := 0
		for _,v := range strs {
			if i == len(v) {
				isBreak = true
				break
			}else if strs[0][:i+1] == v[:i+1] {
				l++
			}else {
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
	}else {
		return ""
	}
}

func main() {
	str := []string {
		//"dog","racecar","car",
		"flower","flow","flowight",
	}
	fmt.Println(longestCommonPrefix(str))
}