package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 二进制转换
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// 一行一行读文件
func printFile(filename string) {
	file, err := os.Open(filename) // 先打开文件 
	if err != nil {
		panic(err)  // 如果有错误，抛出异常
	}

	printFileContents(file) // 读文件
}

// 只有终止条件，相当于 while
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader) // 一行一行读

	for scanner.Scan() {  // go没有while循环， 直接用终止条件的for就是while
		fmt.Println(scanner.Text())
	}
}

// 省略初始，省略结束，这样就是死循环了
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println("convertToBin results:")
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		convertToBin(72387885),
		convertToBin(0),
	)

	fmt.Println("abc.txt contents:")
	printFile("basic/branch/abc.txt")

	fmt.Println("printing a string:")
	s := `abc"d"
	kkkk
	123

	p`
	printFileContents(strings.NewReader(s))

	// Uncomment to see it runs forever
	// forever()
}
