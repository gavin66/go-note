/*
for
Go 只有一种循环结构 for 循环
基本的 for 循环除了没有 () 之外,看起来跟 C 或者 Java 中做的一样,而 {} 是必须的
 */
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
