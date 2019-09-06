package main

import (
	"fmt"
)

/*给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
例如，给出 n = 3，生成结果为：
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

*/

/*回溯法:
回溯法思路的简单描述是：把问题的解空间转化成了图或者树的结构表示，然后使用深度优先搜索策略进行遍历，遍历的过程中记录和寻找所有可行解或者最优解。
基本思想类同于：
   图的深度优先搜索
   二叉树的后序遍历
*/
func generateParenthesis(n int) []string {
	var res []string
	backTrack(&res,"",0,0,n)
	return res
}

func backTrack(res *[]string,cur string,open int,close int,max int) {
	if len(cur) == max *2 {
		fmt.Println("return")
		*res = append(*res,cur)
		return
	}
	if open < max {
		backTrack(res,cur+"(",open+1,close,max)
	}
	if close < open {
		backTrack(res,cur+")",open,close+1,max)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}