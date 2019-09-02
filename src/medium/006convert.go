package main

/*将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
请你实现这个将字符串进行指定行数变换的函数：string convert(string s, int numRows);
示例1:  输入: s = "LEETCODEISHIRING", numRows = 3  输出: "LCIRETOESIIGEDHN"  ;
示例2:  输入: s = "LEETCODEISHIRING", numRows = 4  输出: "LDREOEIIECIHNTSG"  ;
解释:
L     D     R
E   O E   I I
E C   I H   N
T     S     G

------- 这个题比起说是“Z”，还不如说是“N”。（就是找规律）
*/

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	str := ""
	n := len(s)
	cycklLen := 2*numRows-2
	for i := 0;i < numRows;i++ {
		//var arr []string
		for j := 0;j < n - i;j += cycklLen {
			//arr = append(arr,s[j+i:j+i+1])
			str = fmt.Sprint(str,s[j+i:j+i+1])
			if i != 0 && i != numRows-1 && j+cycklLen-i < n {
				//arr = append(arr,s[j+cycklLen-i:j+cycklLen-i+1])
				str = fmt.Sprint(str,s[j+cycklLen-i:j+cycklLen-i+1])
			}
		}
		//fmt.Println(arr)
	}
	return str
}

func main() {
	str := "abcde"
	fmt.Println(convert(str,4))
}