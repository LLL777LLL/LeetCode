package main

import "fmt"

/*报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：
1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。
注意：整数顺序将表示为一个字符串。

示例1:  输入: 1  输出: "1"    ;
示例2:  输入: 4  输出: "1211" ;

*/

//用tmp记录字符串的第一个元素，之后用tmp与后面的作比较，如果第一个元素和后一个不一样，则将第一个储存到新的字符串，将tmp作为还未比较的字符串的第一个元素，以此内推。
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}else {
		str := countAndSay(n-1)
		l := 1
		tmp := str[0:1]
		s := ""
		for i := 0;i < len(str);i++ {
			if i == len(str) - 1 {
				s = fmt.Sprint(s,l,tmp)
			}else {
				if tmp == str[i+1:i+2] {
					l++
				}else {
					s = fmt.Sprint(s,l,tmp)
					tmp = str[i+1:i+2]
					l = 1
				}
			}
		}
		return s
	}
}


//没有将第一个元素设置为比较变量，直接前后比较，太过直观暴力，判断条件也太多。  （不推荐）
func countAndSay2(n int) string {
	if n == 1 {
		return "1"
	}else {
		str := countAndSay2(n-1)
		l := 1
		s := ""
		for i := 0;i < len(str);i++ {
			if len(str[i:]) > 1 {
				if str[i] == str[i+1] {
					l++
				}else {
					if l != 1 {
						s = fmt.Sprint(s,l,str[i:i+1])
						l = 1
					}else {
						s = s + "1" + str[i:i+1]
					}
				}
			}else {
				if l != 1 {
					s = fmt.Sprint(s,l,str[i:i+1])
					l = 1
				}else {
					s = s + "1" + str[i:i+1]
				}
			}
		}
		return s
	}
}

func main() {
	//for i := 1;i < 20;i++ {
	//	fmt.Println(countAndSay(i))
	//}
	fmt.Println(countAndSay(30))
}