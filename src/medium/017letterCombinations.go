package main

import "fmt"

/*给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
示例:  输入："23" ; 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:  尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

*/

//提交了六次，主要是细节上的问题，修改了一个地方之后，有些需要修改的东方没有被修改。
func letterCombinations(digits string) []string {
	var res []string
	isFirst := true
	cons := map[string][]string{
		"2":[]string{"a","b","c"},
		"3":[]string{"d","e","f"},
		"4":[]string{"g","h","i"},
		"5":[]string{"j","k","l"},
		"6":[]string{"m","n","o"},
		"7":[]string{"p","q","r","s"},
		"8":[]string{"t","u","v"},
		"9":[]string{"w","x","y","z"},
	}
	if len(digits) < 2 {
		res = cons[digits]
		return res
	}else {
		for i := 1;i < len(digits);i++ {
			s1 := cons[digits[i:i+1]]
			var s2 []string
			if isFirst {
				res = cons[digits[0:1]]
				isFirst = false
			}
			for j := 0;j < len(res);j++ {
				for k := 0;k < len(s1);k++ {
					s := fmt.Sprint(res[j],s1[k])
					fmt.Println("s:",s)
					s2 = append(s2,s)
				}
			}
			res = s2
		}
	}
	return res
}

func main() {
	str := "2"
	fmt.Println(letterCombinations(str))
}
