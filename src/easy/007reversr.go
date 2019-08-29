package main

import (
	"fmt"
)

/*给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
示例1:  输入: 123    输出: 321;
示例2:  输入: -123   输出: -321;
示例3:  输入: 120    输出: 21 ;
注意:
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。
请根据这个假设，如果反转后整数溢出那么就返回 0。
*/

func reverse(x int) int {
	//用常规的将 数字 倒着排 (注意：将 int 转为 string后，string要用切片的方式分割，如果用 for 遍历的话就会返回 ASCII )
	//var str2 string
	//str := strconv.Itoa(x)
	//fmt.Println(str)
	//l := len(str)
	//if l == 1 {
	//	return x
	//}else {
	//	if str[0] == '-' {
	//		for i := l - 1;i > 0;i-- {
	//			str2 = fmt.Sprint(str2,str[i:i+1])
	//		}
	//		r,_ := strconv.Atoi(str2)
	//		if -r < -2147483648 {
	//			return 0
	//		}
	//		return -r
	//	}else {
	//		for i := l - 1;i >= 0;i-- {
	//			str2 = fmt.Sprint(str2,str[i:i+1])
	//		}
	//		r,_ := strconv.Atoi(str2)
	//		if r > 2147483647 {
	//			return 0
	//		}
	//		return r
	//	}
	//}

	//用数学的方式
	num := 0
	for x != 0 {
		num = num*10 + x % 10
		x /= 10
	}
	if num > 2147483647 || num < -2147483648 {
		return 0
	}
	return num
}

func main(){
	x := -2147
	r := reverse(x)
	fmt.Println(r)

}