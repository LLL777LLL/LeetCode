package main

import (
	"fmt"
	"sort"
)

/*给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。
返回这三个数的和。假定每组输入只存在唯一答案。
例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

*/

//双指针
func threeSumClosest(nums []int,target int) int {
	sort.Ints(nums)
	l := len(nums)
	result,minDistance := 0,0
	isTrue := true
	for i := 0;i < l-2;i++ {
		left := i+1
		right := l-1
		for left < right {
			distance := 0
			sum := nums[i] + nums[left] + nums[right]
			if sum >= target {
				distance = sum - target
			}else {
				distance = target - sum
			}
			if isTrue  {
				result = sum
				minDistance = distance
				isTrue = false
			}else {
				if minDistance > distance {
					result = sum
					minDistance = distance
				}
			}
			fmt.Println(nums[i] , nums[left] ,nums[right],"sum:",sum,"--- dis:",distance," ---- min:",minDistance," ----- i:",i," -----left:",left," --- right:",right)
			if sum >= target {
				right--
			}else {
				left++
			}
		}
	}
	return result
}



//直接三循环，太过于暴力
func threeSumClosest2(nums []int,target int) int {
	l :=len(nums)
	var num int
	var minDistance int
	for i := 0;i < l - 2;i++ {
		j := i +1
		for ;j < l - 1;j++ {
			for k := j + 1;k < l;k++ {
				distance := 0
				sum := nums[i] + nums[j] +nums[k]
				if sum >= target {
					distance = sum - target
				}else {
					distance = target - sum
				}
				if k == 2 {
					minDistance = distance
					num = sum
				}else {
					if minDistance > distance {
						minDistance = distance
						num = sum
					}
				}
				//fmt.Println(nums[i] , nums[j] ,nums[k],"sum:",sum,"--- dis:",distance," ---- min:",minDistance," ----- i:",i," -----j:",j," --- k:",k)
			}
		}
	}
	return num
}

func main() {
	nums := []int{1,-3,3,5,4,1}
	fmt.Println(threeSumClosest(nums,1))
}