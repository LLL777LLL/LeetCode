package main

import "fmt"

/*给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器，且 n 的值至少为 2。
图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
---- 图：就是一个第一象限的坐标图，数组中的元素分别代表Y值，各个元素在数组中的位置代表X值；已数组中第二个元素8为例，其坐标为：（2,8）

示例:  输入: [1,8,6,2,5,4,8,3,7]  输出: 49 ;
解释：坐标（2,8）与坐标（8,7）两个柱子中可以盛的最多。（两个坐标间X轴上的值为7，Y轴上，7为最小，所以输出 49）
*/

/*双指针法:
这种方法背后的思路在于，两线段之间形成的区域总是会受到其中较短那条长度的限制。此外，两线段距离越远，得到的面积就越大。
最初我们考虑由最外围两条线段构成的区域。现在，为了使面积最大化，我们需要考虑更长的两条线段之间的区域。如果我们试图将指向较长线段的指针向内侧移动，矩形
区域的面积将受限于较短的线段而不会获得任何增加。但是，在同样的条件下，移动指向较短线段的指针尽管造成了矩形宽度的减小，但却可能会有助于面积的增大。因为
移动较短线段的指针会得到一条相对较长的线段，这可以克服由宽度减小而引起的面积减小。
*/
func maxArea(height []int) int {
	max,l := 0,0
	r := len(height)-1
	for l < r {
		fmt.Println(l," ------------ r:",r)
		if height[l] < height[r] {
			if max < height[l] * (r-l) {
				max = height[l] * (r-l)
			}
			l++
		}else {
			if max < height[r] * (r-l) {
				max = height[r] * (r-l)
			}
			r--
		}
	}
	return max
}

//直接的暴力解法
func maxArea2(height []int) int {
	area := 0
	for i := 0;i < len(height);i++{
		for j := i+1;j < len(height);j++{
			if height[i] > height[j] {
				if area < height[j] * (j-i) {
					area = height[j] * (j-i)
				}
			}else {
				if area < height[i] * (j-i) {
					area = height[i] * (j-i)
				}
			}
		}
	}
	return area
}

func main() {
	height := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(height))
}