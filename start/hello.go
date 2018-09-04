package main

//import "./chen"  import当前目录里chen子目录里的所有的go文件
import "fmt" //import 环境变量 $GOPATH/src/fmt子目录里的所有的go文件
func main() {
	// { 中括号分行，编译报错 与自动格式化加；有关
	fmt.Println("Hello, World")
}

/*
可编译、可解释
1. go run hello.go
2. go build hello.go
   ./hello
*/
