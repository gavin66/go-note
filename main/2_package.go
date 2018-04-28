/*
每个 Go 程序都是由包组成的
程序运行的入口是包 main
当前的程序使用并导入了包"fmt"和"math/rand"
按照惯例,包名与导入路径的最后一个目录一致.例如,"math/rand"包由 package rand 语句开始
 */
package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))
}
