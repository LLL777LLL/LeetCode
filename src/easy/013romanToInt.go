package main

import "fmt"

/*罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，
所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

示例1:  输入: "III"      输出: 3    ;
示例2:  输入: "IV"       输出: 4    ;
示例3:  输入: "IX"       输出: 9    ;
示例4:  输入: "LVIII"    输出: 58   ; 解释: L = 50, V= 5, III = 3.
示例5:  输入: "MCMXCIV"  输出: 1994 ; 解释: M = 1000, CM = 900, XC = 90, IV = 4.

*/


/*仔细观察发现 罗马数字的写法是从大到小的（从从左到右），所以，如果一个比较小的数字在较大的左边的话，那就说明这两个是可以组合成
那六中特殊情况的一种，在那六中特殊情况之外，不可能有较小的在较大的左边的情况发生。
 */
func romanToInt(s string ) int {
	data := map[string]int {
		"I":1,
		"V":5,
		"X":10,
		"L":50,
		"C":100,
		"D":500,
		"M":1000,
	}

	l := len(s)
	i,sum := 0,0
	for i < l {
		if l-i > 1 && data[s[i:i+1]] < data[s[i+1:i+2]] {
			sum +=  data[s[i+1:i+2]] - data[s[i:i+1]]
			i += 2
			fmt.Println(sum)
		}else {
			sum += data[s[i:i+1]]
			i += 1
			fmt.Println(" 2 ----------- ",sum)
		}
	}
	return sum
}

//这是最常规的方式
func romanToInt2(s string) int {
	data := map[string]int {
		"I":1,
		"IV":4,
		"IX":9,
		"V":5,
		"X":10,
		"XL":40,
		"XC":90,
		"L":50,
		"C":100,
		"CD":400,
		"CM":900,
		"D":500,
		"M":1000,
	}

	l := len(s)
	sum := 0
	for i:= 0;i < l;i++ {
		if len(s) < 1 {
			break
		}
		isGo := false
		for k,v := range data {
			if len(s) < 2 {
				break
			}
			if s[:2] == k {
				sum += v
				s = s[2:]
				isGo = true
				break
			}
		}
		if !isGo {
			for k,v := range data {
				if s[:1] == k {
					sum += v
					s = s[1:]
					break
				}
			}
		}
	}
	return sum
}

func main() {
	str := "MCMXCIV"
	fmt.Println(romanToInt(str))
}