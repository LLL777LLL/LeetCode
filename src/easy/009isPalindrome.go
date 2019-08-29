package main

import "fmt"

/*判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
示例1:  输入: 121  输出: true ;
示例2:  输入: -121  输出: false ;  解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例3: 输入: 10  输出: false ;     解释: 从右向左读, 为 01 。因此它不是一个回文数。

*/

func isPalindrome(x int) bool {
	num := x
	if num < 0 {
		return false
	}
	flag := 0
	for num != 0 {
		flag = flag * 10 + num % 10
		num /= 10
	}
	return flag == x
}

func main() {
	x := 123321
	fmt.Println(isPalindrome(x))
}