package main

import (
	"fmt"
)

/*实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须原地修改，只允许使用额外常数空间。
以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

*/

func nextPermutaton(nums []int) {
	l := len(nums)
	i := l-2
	if l < 2 {
		return
	}
	isTrue := true
	for i >= 0 && nums[i+1] <= nums[i] {
		fmt.Println("for 1 ------- i:",i)
		i--
	}
	if i >= 0 {
		j := i+1
		for ;j < l;j++ {
			if nums[i] >= nums[j] {
				fmt.Println("for 2 ------- i:",i," ------------ j:",j)
				nums[i],nums[j-1] = nums[j-1],nums[i]
				isTrue = false
				break
			}
		}
		if isTrue {
			//解决类似 1,2,3,4,5,5,5,5,5,5  这样末尾为重复的数字的情况.(与上面变量为j的for循环冲突，意思是，上面个for执行了break，那此时就用执行，否则就执行。)
			nums[i],nums[j-1] = nums[j-1],nums[i]
		}
	}
	j := l-1
	for i < j {
		i++
		nums[i],nums[j] = nums[j],nums[i]
		j--
	}
}

func nextPermutaton2(nums []int) {
	l := len(nums)
	if l < 2 {
		return
	}
	i := l-2
	for i >= 0 && nums[i-1] < nums[i] {
		i--
	}
	if i >= 0 {
		//j := l+1

	}
}

func main() {
	arr := []int{1,1,5}
	nextPermutaton(arr)
	fmt.Println(arr)
}