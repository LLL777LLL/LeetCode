package main

import (
	"fmt"
	"strings"
)

/*实现 strStr() 函数。
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
示例1:  输入: haystack = "hello", needle = "ll"    输出: 2  ;
示例2:  输入: haystack = "aaaaa", needle = "bba"   输出: -1 ;

说明:
当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/



func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	nel := len(needle)
	for i := 0;i < len(haystack);i++ {
		if len(haystack) - i >= nel {
			if haystack[i:i+nel] == needle {
				return i
			}
		}
	}
	return -1
}


//使用库函数 strings
func strStr2(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	i := strings.Index(haystack,needle)
	return i
}



func main() {
	str := "abcdef"
	str2 := "g"
	i := strStr(str,str2)
	fmt.Println(i)
}