package rightShift

/*
场景应用：
有时候在项目中可能会用到同类别的不同分类，给不同分类定义一个唯一识别字段，比如：
用 << 左移生成唯一识别号，1,2,4，8,16 。。。。，之后类别是不同分类的和 ， 如：1+4+8 = 13
所以 拿到 13 后，需要解析出来 其中有哪些类。
*/

import (
	"math"
)

func Right(num int)[]int {
	var (
		arr1 []int
		arr []int
		count int
	)

	//将十进制转换成二进制（用arr1 切片存储，不过是二进制结果的倒序，便于之后用 range 的 k 将其转换成 各个数）
	for ;num > 0;num /= 2 {
		count = num % 2
		arr1 = append(arr1,count)
	}

	//因为存储的时候是二进制结果倒序存储的，所以遍历的索引正好的 2 的幂
	for k,v := range arr1 {
		if v == 1 {
			numb := math.Pow(2,float64(k))
			arr = append(arr,int(numb))
		}
	}
	return arr
}

func HasRight(num1,num2 int) bool {
	arr := Right(num2)
	for _,v := range arr {
		if v == num1 {
			return true
		}
	}
	return false
}

