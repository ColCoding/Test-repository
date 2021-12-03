package main

import "fmt"

func main() {
	var n [10]int /* n 是一个长度为 10 的数组 */

	/* 为数组 n 初始化元素 */
	for i := 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for i := 0; i < 10; i++ {
		fmt.Printf("Element[%d] = %d\n", i, n[i])
	}

}
