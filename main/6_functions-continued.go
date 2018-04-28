/*
函数(续)
当两个或多个连续的函数命名参数是同一类型,则除了最后一个类型之外,其他都可以忽略
在这个例子中
x int, y int
被缩写为
x, y int
 */
package main

import "fmt"

func add2(x, y int) int {
	return x+y
}

func main() {
	fmt.Println(add2(3,33))
}