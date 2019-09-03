package main

import (
	"fmt"
	"sort"
)

/*给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/


//------------ 提交的第六次才通过
/****** 错误统计：
1、执行出错信息：
	Line 5: panic: runtime error: index out of range
	最后执行的输入：[]
2、执行出错信息：
	Line 24: panic: runtime error: index out of range
	最后执行的输入：[0,0,0]
3、输入：[0,0,0,0] ；   输出：[[0,0,0],[0,0,0]]   ； 预期：[[0,0,0]]；
4、执行出错信息：
	Line 41: panic: runtime error: index out of range
	最后执行的输入：[4,4,-1,4,-1]
5、输入：[-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6]
	输出：[[-4,-2,6],[-4,0,4],[-4,1,3],[-4,2,2]]
	预期：[[-4,-2,6],[-4,0,4],[-4,1,3],[-4,2,2],[-2,-2,4],[-2,0,2]]
*/
func threeSum(nums []int) [][]int {
	l := len(nums)
	var res [][]int
	if l < 3 {
		return res
	}
	sort.Ints(nums)
	fmt.Println(nums)
	if nums[0] <= 0 && nums[l-1] >= 0 {  //保证有正负数
		i := 0
		for i < l - 2 {
			left := i+1
			last := l - 1
			for left < last {
				if nums[i] * nums[last] > 0 {  //三数同号，无解
					break
				}
				sum := nums[i] + nums[left] + nums[last]
				if sum == 0 {
					num := []int{nums[i],nums[left],nums[last]}
					res = append(res,num)
				}
				if sum <= 0 {
					left++
					for nums[left-1] == nums[left] {  //去重
						if left >= last {
							break
						}
						left++
					}
				}else {
					last--
					for nums[last] == nums[last+1] {  //去重
						if last <= left {
							break
						}
						last--
					}
				}
			}
			i++
			for nums[i-1] == nums[i] {  //去重
				if i >= l - 2 {
					break
				}
				i++
			}
		}
	}
	return res
}

func main() {
	nums := []int{0,0,0,0,0,0,0,0,0}
	fmt.Println(threeSum(nums))
}