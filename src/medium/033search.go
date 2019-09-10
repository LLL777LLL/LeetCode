package main

import (
	"fmt"
)

/*假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
你可以假设数组中不存在重复的元素。
你的算法时间复杂度必须是 O(log n) 级别。

示例1:  输入: nums = [4,5,6,7,0,1,2], target = 0  输出: 4   ;
示例2:  输入: nums = [4,5,6,7,0,1,2], target = 3  输出: -1  ;

*/

//先用二分法找到旋转点下标 index ，之后就好做了。
func search(nums []int,target int) int {
	l := len(nums)
	if l == 1{
		if nums[0] == target {
			return 0
		}
		return -1
	}
	if l == 0 {
		return -1
	}
	left,right := 0,l-1
	for right - left != 1 {
		middle :=  (left+right) / 2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] < nums[right] {
			right = middle
		}else {
			left = middle
		}
	}

	//此时nums中最大值得下标为left，最小值下标right;
	l,r := 0,l-1
	if target <= nums[r] && target >= nums[right] {
		l = right
	}else {
		r = left
	}
	for l <= r {
		mid := (l+r) / 2
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			l = mid+1
		}else {
			r = mid-1
		}
	}
	return -1
}

func main() {
	nums := []int{6}
	fmt.Println(search(nums,6))
}
