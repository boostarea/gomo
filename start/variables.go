package main

import "fmt"

/**
initial
1 2

short
*/
func main() {

	// 初始化后未使用，则编译报错
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	//var d int
	var e string
	fmt.Println(e)

	// 仅用于局部变量
	f := "short"
	fmt.Println(f)
}
