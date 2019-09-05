package main

import (
	"fmt"
	"sort"
)

/*给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值
与 target 相等？找出所有满足条件且不重复的四元组。
注意：答案中不可以包含重复的四元组。

示例：给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

------ 未通过情况：
1、-3,-2,-1,0,0,1,2,3 ; 0
2、0,0,0,0 ;  0
3、-1,0,1,2,-1,-4 ;-1
4、-1,-5,-5,-3,2,5,0,4  ; -7
*/

//双指针法 -- 先排序（升序），固定前两个数，之后用两个指针指向头和尾，四个数之和 sum 大于 target，则尾左移，否则头右移
// -------------- 提交的第七次才通过，很多情况都没有考虑到！！！！！！！！！
func fourSum(nums []int,target int) [][]int {
	l := len(nums)
	var res [][]int
	if l < 4 {
		return res
	}
	sort.Ints(nums)
	var sum int
	for i := 0;i < l-3;i++ {
		if i >0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i+1;j < l-2;j++ {
			if j >1 && i != j-1 && nums[j-1] == nums[j] {
				continue
			}
			left := j+1
			right := l-1
			for left < right {
				sum = nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					fmt.Println("i:",i," ------ j:",j," ---- left:",left," ------------- right:",right," ---------sum:",sum)
					s := []int{nums[i],nums[j],nums[left],nums[right]}
					res = append(res,s)
				}
				if sum >= target {
					right--
					for nums[right+1] == nums[right] && right > left {
						right--
					}
				}else {
					left++
					for nums[left-1] == nums[left]  && left < right {
						left++
					}
				}
			}
		}
	}
	return res
}

func main() {
	nums := []int{0,0,0,0,0,0,0,1}
	fmt.Println(fourSum(nums,0))
	fmt.Println(nums)
}