package main

import "fmt"

// slice的长度和容量是多少呢？系统又是怎么分配的呢？
func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n",
		s, len(s), cap(s))
}

// 通过打印发现，len是递增的，cap是2^n分配的，直到装下的最小容量；
func sliceOps() {
	fmt.Println("Creating slice")// slice的创建，go有个零值，叫nil,没有null;
	var s []int // Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)


	// 创建slice的另外一个钟方法
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	// 通过make来创建slice,直到大小，不知道值；可以指定len，和cap
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	// slice的拷贝
	fmt.Println("Copying slice")
	copy(s2, s1) // 系统的copy函数
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	// 删除中间一个数，因为没有系统函数，通过slice的append方法可变参数移除；
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	printSlice(s2)
}
