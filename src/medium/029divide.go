package main

import "fmt"

/*给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
返回被除数 dividend 除以除数 divisor 得到的商。
示例1:  输入: dividend = 10, divisor = 3  输出: 3   ;
示例2:  输入: dividend = 7, divisor = -3  输出: -2  ;

说明:
被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。(2的31次方)

*/

//使用二分法  (特别要注意的是，溢出点：2147483647，-2147483648)
func divide(dividend ,divisor int) int {
	var res int
	if dividend == 0 {
		return res
	}
	if divisor == 1 || divisor == -1 {
		if dividend * divisor > 2147483647 {
			return 2147483647
		}else if dividend * divisor < -2147483648 {
			return -2147483648
		}else {
			return dividend * divisor
		}
	}
	flag1,flag2 := 1,1
	if dividend < 0 {
		dividend = -dividend
		flag1 = -1
	}
	if divisor < 0 {
		divisor = -divisor
		flag2 = -1
	}
	start,end := 0,dividend
	middle := (start+end) / 2
	for start < end {
		if middle * divisor > dividend {
			end = middle
			middle = (start+end) / 2
		}else if (middle+1) * divisor >= dividend {
			num := flag1*flag2*middle
			if num < -2147483648 {
				res = -2147483648
			}else if num > 2147483647 {
				res = 2147483647
			}else {
				res = num
			}
			break
		}else {
			start = middle
			middle = (start+end) / 2
		}
	}
	return res
}

func main() {
	fmt.Println(divide(2,2))
}
