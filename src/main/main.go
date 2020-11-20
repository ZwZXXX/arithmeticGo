package main

import (
	"arithmetic/src/sort"
	"fmt"
)

/**
* @Author: ZwZ
* @Time : 2020/11/17 15:31
* @File : main
* @Description:
 */
func main() {
	var arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.BubbleSort1(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
