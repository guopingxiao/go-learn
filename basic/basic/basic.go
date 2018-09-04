package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/* 可以通过括号的形式来避免多次var */
var (
	aa = 3
	ss = "kkk"
	bb = true
)
/*
	初始化变量
*/
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}
/*
	初始化并赋值
*/
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

/*
	可以不声明类型
*/
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}
/*
	函数里初始化赋值的另一只方式,自动类型判断
*/
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}
// 欧拉函数 1i表示复数
func euler() {
	a:=3+4i
	fmt.Println(cmplx.Abs(a))
	
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

// 勾股定理
func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}
// 数据的强制类型转换
func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

// 常量 const
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//  枚举值，直接用const来定义，有两类，一类是普通的，一类是自增的枚举值，iota
func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()
	consts()
	enums()
}
