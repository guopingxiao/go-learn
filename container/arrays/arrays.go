package main

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int // 只赋值，不初始化
	arr2 := [3]int{1, 3, 5} // :=的形式，要给定个数和初始值
	arr3 := [...]int{2, 4, 6, 8, 10} // :=的形式，[...]可以让系统给你数个数
	var grid [4][5]int // 二维数组

	fmt.Println("array definitions:")
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	fmt.Println("printArray(arr1)")
	printArray(arr1)

	fmt.Println("printArray(arr3)")
	printArray(arr3)

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
}
